package registration

import (
	"github.com/google/uuid"
)

// Register places a new server into the registry, and returns an API key for
// the server.
func Register(name string, userID string) (string, error) {
	key := generateKey()

	// TODO store key in database

	return key, nil
}

func generateKey() string {
	// TODO crypto safe key generation
	return uuid.New().String()
}
