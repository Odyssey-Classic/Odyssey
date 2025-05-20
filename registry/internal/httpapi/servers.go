package httpapi

import (
	"github.com/FosteredGames/Odyssey/registry/internal/registry/servers"
	"github.com/go-chi/chi/v5"
)

// ServersAPI returns a chi.Router for all /servers endpoints.
func ServersAPI(service *servers.Service) chi.Router {
	r := chi.NewRouter()
	handlers := NewServersHandlers(service)

	r.Post("/register", handlers.Register)
	// Add more endpoints here as needed

	return r
}
