package server

import (
	"log/slog"
	"net/http"

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
	key := r.Header.Get("Authorization")
	return h.servers.FindByKey(r.Context(), key)
}
