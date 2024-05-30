package identity

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/FosteredGames/Odyssey/registry/internal/config"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity/oauth"
	"github.com/golang-jwt/jwt/v5"
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

	oAuthServer := &oauth.OAuthServer{
		Config:           idServer.oAuth,
		IdentityCallback: idServer.IdentityCallback,
	}
	mux.HandleFunc("/login", oAuthServer.OAuthRedirect)

	return idServer
}

func (s *IdentityServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("identity server")
	s.mux.ServeHTTP(w, r)
}

func (s *IdentityServer) IdentityCallback(ctx context.Context, id string) (string, error) {
	fmt.Printf("identity callback: %v\n", id)
	// if s.users[id] == "" {
	s.newUser(ctx, id)
	fmt.Printf("new user: %v\n", id)
	tok, err := s.GenerateJWT(id)
	if err != nil {
		fmt.Printf("failed to generate jwt: %v\n", err)
		return "", err
	}
	// s.users[id] = tok
	// }

	fmt.Printf("jwt %s\n", tok)

	return tok, nil
}

func (s *IdentityServer) newUser(ctx context.Context, id string) {
	db := s.db.Client.Database("registry").Collection("users")
	user := User{
		DiscordID: id,
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := db.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
	}
}

func (s *IdentityServer) GenerateJWT(id string) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": id,
	})

	return tok.SignedString(s.privateKey)
}
