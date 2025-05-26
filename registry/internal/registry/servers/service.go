package servers

import (
	"context"
	"log/slog"
	"sync"

	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// RegisterServer registers a new game server.
func (s *Service) RegisterServer(ctx context.Context, name string, userID string) (APIKey, error) {
	db := s.db.Client.Database("registry").Collection("servers")

	key, hash, err := generateKey()
	if err != nil {
		return "", err
	}

	server := ServerInfo{
		Key:  hash,
		Name: name,
		User: primitive.ObjectID.MarshalText(userID),
	}

	res, err := db.InsertOne(ctx, server, &options.InsertOneOptions{Comment: "registring new server"})
	if err != nil {
		return "", err
	}

	slog.Info("[servers] new server registered", "id", res.InsertedID)

	return key, nil
}

// ListServers returns all registered servers.
func (s *Service) ListServers() []ServerInfo {
	s.mu.Lock()
	defer s.mu.Unlock()
	servers := make([]ServerInfo, 0, len(s.servers))
	for _, srv := range s.servers {
		servers = append(servers, srv)
	}
	return servers
}
