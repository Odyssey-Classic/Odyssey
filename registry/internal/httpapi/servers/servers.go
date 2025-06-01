package servers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/FosteredGames/Odyssey/registry/internal/httpapi/identity"
	users "github.com/FosteredGames/Odyssey/registry/internal/registry/identity"
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

	r.Post("/", api.register)
	r.Get("/", api.listServers)
	r.Get("/mine", api.listUserServers)

	return api
}

// Register handles POST /register for game server registration.
func (h *API) register(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(identity.UserKeyContext).(*users.User)
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

	key, err := h.servers.RegisterServer(r.Context(), info.Name, user)
	if errors.Is(err, servers.ErrServerLimitReached) {
		http.Error(w, "max number of servers reached", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "registration failed", http.StatusInternalServerError)
		slog.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "registered", "apiKey": string(key)})
}

func (a *API) listServers(w http.ResponseWriter, r *http.Request) {
	list, err := a.servers.ListServers(r.Context())
	if err != nil {
		http.Error(w, "failed to list servers", http.StatusInternalServerError)
		return
	}

	jsonList := make([]map[string]string, len(list))
	for i, server := range list {
		jsonList[i] = map[string]string{"id": server.ID.Hex(), "name": server.Name, "owner": server.User.Hex()}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonList)
}

func (a *API) listUserServers(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(identity.UserKeyContext).(*users.User)
	if !ok {
		http.Error(w, "unauthenticated", http.StatusForbidden)
		slog.Error("[servers] user info missing from context")
		return
	}

	list, err := a.servers.ListUserServers(r.Context(), user)
	if err != nil {
		http.Error(w, "failed to list servers", http.StatusInternalServerError)
		return
	}

	jsonList := make([]map[string]string, len(list))
	for i, server := range list {
		jsonList[i] = map[string]string{"id": server.ID.Hex(), "name": server.Name, "owner": server.User.Hex()}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonList)
}
