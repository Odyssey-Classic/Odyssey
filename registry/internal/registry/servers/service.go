package servers

import (
	"context"
	"errors"
	"log/slog"
	"sync"

	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Service handles business logic for server registration and management.
type Service struct {
	mu      sync.Mutex
	servers map[string]ServerInfo
	db      *data.DB
}

func NewService(db *data.DB) *Service {
	return &Service{
		servers: make(map[string]ServerInfo),
		db:      db,
	}
}

var ErrServerLimitReached = errors.New("user has reached the server limit")

// RegisterServer registers a new game server.
func (s *Service) RegisterServer(ctx context.Context, name string, user *identity.User) (APIKey, error) {
	db := s.db.Client.Database("registry").Collection("servers")

	var result bson.M
	err := db.FindOne(ctx, bson.D{{Key: "user", Value: user.ID}}).Decode(&result)

	// If we find one, or there's an error.
	// Naive way to limit users to one server.
	if !errors.Is(err, mongo.ErrNoDocuments) {
		if err == nil {
			err = ErrServerLimitReached
		}
		return "", err
	}

	key, hash, err := generateKey()
	if err != nil {
		return "", err
	}

	server := ServerInfo{
		Key:  hash,
		Name: name,
		User: user.ID,
	}

	res, err := db.InsertOne(ctx, server, &options.InsertOneOptions{Comment: "registring new server"})
	if err != nil {
		return "", err
	}

	slog.Info("[servers] new server registered", "id", res.InsertedID)

	return key, nil
}

// ListServers returns all registered servers.
func (s *Service) ListServers(ctx context.Context) ([]ServerInfo, error) {
	return s.getServers(ctx, bson.D{})
}

func (s *Service) ListUserServers(ctx context.Context, user *identity.User) ([]ServerInfo, error) {
	return s.getServers(ctx, bson.D{{Key: "user", Value: user.ID}})
}

func (s *Service) FindByKey(ctx context.Context, key string) ([]ServerInfo, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return s.getServers(ctx, bson.D{{Key: "Key", Value: hash}})
}

func (s *Service) getServers(ctx context.Context, filter bson.D) ([]ServerInfo, error) {
	db := s.db.Client.Database("registry").Collection("servers")
	cur, err := db.Find(ctx, filter)
	if err != nil {
		slog.Error("[servers] failed to list servers", "error", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var servers []ServerInfo
	for cur.Next(ctx) {
		var server ServerInfo
		if err := cur.Decode(&server); err != nil {
			slog.Error("[servers] failed to decode server", "error", err)
			continue
		}
		servers = append(servers, server)
	}
	if err := cur.Err(); err != nil {
		slog.Error("[servers] cursor error", "error", err)
		return nil, err
	}

	return servers, nil
}
