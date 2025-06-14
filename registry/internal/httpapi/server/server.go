package server

import (
	"encoding/base64"
	"net/http"
	"strings"

	"errors"
	"log/slog"

	"github.com/FosteredGames/Odyssey/registry/internal/registry/servers"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

var ErrUnauthorized = errors.New("unauthorized")

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

	r.Get("/", api.ping)

	return api
}

func (h *API) ping(w http.ResponseWriter, r *http.Request) {
	info, err := h.getServer(r)
	if err != nil || len(info) <= 0 {
		http.Error(w, "unable to find server", http.StatusNotFound)
		slog.Error("[server] error finding server", "error", err)
		return
	}

	slog.Info("[server] ping", "server", info[0])
}

func (h *API) getServer(r *http.Request) ([]servers.ServerInfo, error) {
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Basic ") {
		return nil, ErrUnauthorized
	}
	payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
	if err != nil {
		return nil, err
	}
	parts := strings.SplitN(string(payload), ":", 2)
	if len(parts) != 2 {
		return nil, ErrUnauthorized
	}
	id, key := parts[0], parts[1]
	serversList, err := h.servers.FindByID(r.Context(), id)
	if err != nil || len(serversList) == 0 {
		return nil, ErrUnauthorized
	}
	server := serversList[0]
	if err := bcrypt.CompareHashAndPassword([]byte(server.Key), []byte(key)); err != nil {
		return nil, ErrUnauthorized
	}
	return []servers.ServerInfo{server}, nil
}
