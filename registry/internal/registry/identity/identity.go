package identity

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/FosteredGames/Odyssey/registry/internal/data"
	"github.com/FosteredGames/Odyssey/registry/internal/identity/oauth"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

// IdentityServer is the root HTTP server for basic identity operations.
type IdentityServer struct {
	db *data.DB

	oAuth      *oauth2.Config
	privateKey *ecdsa.PrivateKey

	users map[string]string
}

func New(privateKey *ecdsa.PrivateKey, oAuth *oauth2.Config, db *data.DB) *IdentityServer {
	return &IdentityServer{
		db:         db,
		oAuth:      oAuth,
		privateKey: privateKey,
		users:      make(map[string]string, 1),
	}
}

func (s *IdentityServer) Handler() http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", s.Login)
	mux.HandleFunc("/oauth/callback", s.OAuthCallback)

	return mux
}

func (s *IdentityServer) Run(ctx context.Context) {
	oAuthServer := &oauth.OAuthServer{
		Config:           s.oAuth,
		IdentityCallback: s.IdentityCallback,
	}

	oAuthServer.Run()
}

func (s *IdentityServer) IdentityCallback(ctx context.Context, id string) (string, error) {
	fmt.Printf("identity callback: %v\n", id)
	if s.users[id] == "" {
		s.newUser(ctx, id)
		fmt.Printf("new user: %v\n", id)
		tok, err := s.GenerateJWT(id)
		if err != nil {
			fmt.Printf("failed to generate jwt: %v\n", err)
			return "", err
		}
		s.users[id] = tok
	}

	fmt.Printf("jwt %s\n", s.users[id])

	return s.users[id], nil
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

func (s *IdentityServer) VerifyJWT(token string) (string, error) {
	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return s.privateKey.Public(), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to parse claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse sub")
	}

	return sub, nil
}
