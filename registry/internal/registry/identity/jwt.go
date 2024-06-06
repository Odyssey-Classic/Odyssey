package identity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type UserKey string

const UserKeyContext UserKey = "jwt-user"

func (s *IdentityServer) verifyJWT(token string) (string, error) {
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

	// TODO: do we need to validate the sub against the database?

	return sub, nil
}

func (s *IdentityServer) AuthorizeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "no token", http.StatusUnauthorized)
			return
		}

		sub, err := s.verifyJWT(token)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// Add the known user to the context
		r.WithContext(context.WithValue(r.Context(), UserKeyContext, sub))
		handler.ServeHTTP(w, r)
	})
}
