package registry

import (
	"context"
	"crypto/ecdsa"
	"net/http"

	"github.com/FosteredGames/Odyssey/registry/internal/config"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/servers"
)

type Registry struct {
	DB          *data.DB
	OAuthConfig config.OAuthConfig
	PrivateKey  *ecdsa.PrivateKey
}

func (r *Registry) Run(ctx context.Context) error {
	mux := http.NewServeMux()

	idServer := identity.New(r.PrivateKey, r.OAuthConfig, r.DB)
	mux.Handle("/identity", http.StripPrefix("/identity", idServer))

	seversServer := &servers.ServersServer{}
	mux.Handle("/servers", idServer.AuthorizeMiddleware(seversServer))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		server.Shutdown(ctx) // TODO: do we need a different context?
	}()

	return server.ListenAndServe()
}
