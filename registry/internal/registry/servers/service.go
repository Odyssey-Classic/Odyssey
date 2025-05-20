package servers

import "sync"

// Service handles business logic for server registration and management.
type Service struct {
	mu      sync.Mutex
	servers map[string]ServerInfo
}

func NewService() *Service {
	return &Service{
		servers: make(map[string]ServerInfo),
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
