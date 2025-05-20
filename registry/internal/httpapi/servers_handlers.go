package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/FosteredGames/Odyssey/registry/internal/registry/servers"
)

type ServersHandlers struct {
	Service *servers.Service
}

func NewServersHandlers(service *servers.Service) *ServersHandlers {
	return &ServersHandlers{Service: service}
}

// Register handles POST /register for game server registration.
func (h *ServersHandlers) Register(w http.ResponseWriter, r *http.Request) {
	var info servers.ServerInfo
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if info.ID == "" || info.Name == "" || info.Address == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	h.Service.RegisterServer(info)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "registered"})
}
