package identity

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/FosteredGames/Odyssey/registry/internal/config"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity/oauth"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
)

// IdentityServer is the root HTTP server for basic identity operations.
type IdentityServer struct {
	db         *data.DB
	oAuth      *oauth2.Config
	privateKey *ecdsa.PrivateKey

	mux *http.ServeMux
}

func New(privateKey *ecdsa.PrivateKey, oAuth config.OAuthConfig, db *data.DB) *IdentityServer {
	mux := http.NewServeMux()

	idServer := &IdentityServer{
		db: db,
		oAuth: &oauth2.Config{
			ClientID:     oAuth.ClientID,
			ClientSecret: oAuth.ClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  oAuth.AuthorizationURL.String(),
				TokenURL: oAuth.TokenURL.String(),
			},
			RedirectURL: oAuth.RedirectURL.String(),
			Scopes:      []string{},
		},
		privateKey: privateKey,
		mux:        mux,
	}

	oAuthServer := oauth.New(idServer.oAuth, idServer.IdentityCallback)

	mux.HandleFunc("/login", oAuthServer.OAuthRedirect)
	mux.HandleFunc("/oauth/callback", oAuthServer.OAuthCallback)

	return idServer
}

func (s *IdentityServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.InfoContext(r.Context(), "Identity Server", "path", r.URL.Path)
	s.mux.ServeHTTP(w, r)
}

func (s *IdentityServer) IdentityCallback(ctx context.Context, id string) (string, error) {
	s.newUser(ctx, id)

	tok, err := s.GenerateJWT(id)
	if err != nil {
		fmt.Printf("failed to generate jwt: %v\n", err)
		return "", err
	}

	return tok, nil
}

func (s *IdentityServer) newUser(ctx context.Context, id string) {
	db := s.db.Client.Database("registry").Collection("users")
	user := User{
		DiscordID: id,
	}

	filter := bson.M{"discord_id": id}

	result, err := db.ReplaceOne(ctx, filter, user, options.Replace().SetUpsert(true))
	_ = result
	if err != nil {
		slog.Error(err.Error())
	}
}

func (s *IdentityServer) GenerateJWT(id string) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": id,
	})

	return tok.SignedString(s.privateKey)
}
