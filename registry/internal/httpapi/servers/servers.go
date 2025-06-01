package servers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/FosteredGames/Odyssey/registry/internal/httpapi/identity"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/servers"
	"github.com/go-chi/chi/v5"
)

type API struct {
	servers *servers.Service
	router  chi.Router
}

func (a *API) Router() chi.Router {
	return a.router
}

func New(servers *servers.Service) *API {
	r := chi.NewRouter()

	api := &API{
		servers: servers,
		router:  r,
	}

	r.Post("/register", api.register)

	return api
}

// Register handles POST /register for game server registration.
func (h *API) register(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(identity.UserKeyContext).(string)
	if !ok {
		http.Error(w, "unauthenticated", http.StatusForbidden)
		slog.Error("[servers] user info missing from context")
		return
	}

	var info servers.ServerInfo
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if info.Name == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	h.servers.RegisterServer(r.Context(), info.Name, user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "registered"})
}
