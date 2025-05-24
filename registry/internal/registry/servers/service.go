package servers

import (
	"sync"

	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
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
func (s *Service) RegisterServer(info ServerInfo) {
	s.mu.Lock()
	s.servers[info.ID] = info
	s.mu.Unlock()
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
