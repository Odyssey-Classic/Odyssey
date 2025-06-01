package identity

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserKey string

const UserKeyContext UserKey = "jwt-user"

func (s *Identity) VerifyJWT(token string) (string, error) {
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
