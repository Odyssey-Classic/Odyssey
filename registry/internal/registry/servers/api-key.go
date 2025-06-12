package servers

import (
	"crypto/rand"
	"encoding/base32"

	"golang.org/x/crypto/bcrypt"
)

type APIKey string

func generateKey() (APIKey, string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", "", err
	}
	key := APIKey(base32.StdEncoding.EncodeToString(b))

	hash, err := hashKey(key)
	if err != nil {
		return "", "", err
	}

	return key, string(hash), nil
}

func hashKey(key APIKey) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
