package registry

import (
	"context"
	"net/http"

	"github.com/FosteredGames/Odyssey/registry/internal/config"
	"github.com/FosteredGames/Odyssey/registry/internal/data"
)

type Registry struct {
	DB *data.DB

	OAuth *config.OAuthConfig
}

func (r *Registry) Run(ctx context.Context) {
	mux := http.NewServeMux()
	mux.Handle("/login", handler http.Handler)
}
