package identity

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/FosteredGames/Odyssey/registry/internal/identity/oauth"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

// IdentityServer is the root HTTP server for basic identity operations.
type IdentityServer struct {
	oAuth      *oauth2.Config
	privateKey *ecdsa.PrivateKey

	users map[string]string
}

func New(privateKey *ecdsa.PrivateKey, oAuth *oauth2.Config) *IdentityServer {
	return &IdentityServer{
		oAuth:      oAuth,
		privateKey: privateKey,
		users:      make(map[string]string, 1),
	}
}

func (s *IdentityServer) Run(ctx context.Context) {
	oAuthServer := &oauth.OAuthServer{
		Config:           s.oAuth,
		IdentityCallback: s.IdentityCallback,
	}

	oAuthServer.Run()
}

func (s *IdentityServer) IdentityCallback(id string) (string, error) {
	fmt.Printf("identity callback: %v\n", id)
	if s.users[id] == "" {
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
