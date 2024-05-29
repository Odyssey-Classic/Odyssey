package registry

import (
	"context"
	"net/http"

	"github.com/FosteredGames/Odyssey/registry/internal/config"
	"github.com/FosteredGames/Odyssey/registry/internal/data"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity"
)

type Registry struct {
	DB *data.DB

	OAuth *config.OAuthConfig
}

func (r *Registry) Run(ctx context.Context) {
	idServer := identity.New(r.OAuth.PrivateKey, r.OAuth.Config, r.DB)

	mux := http.NewServeMux()
	mux.Handle("/identity", handler http.Handler)
}
