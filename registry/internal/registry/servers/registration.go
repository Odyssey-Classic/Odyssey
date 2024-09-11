package servers

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type LimitError struct {
	error
}

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

func (service *Service) findServersByUser(ctx context.Context, userID string) ([]Server, error) {
	db := service.db.Client.Database("registry").Collection("servers")
	db.Find(ctx, bson.M{"registration": bson.M{"userID": userID}})
	// TODO implement
	return nil, nil
}
