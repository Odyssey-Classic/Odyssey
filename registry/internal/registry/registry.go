package registry

import (
	"context"
	"crypto/ecdsa"
	"log/slog"
	"net/http"
	"time"

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
	mux.Handle("/identity/", http.StripPrefix("/identity", idServer))

	serversServer := &servers.ServersServer{}
	mux.Handle("/servers", idServer.AuthorizeMiddleware(serversServer))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		<-ctx.Done()

		slog.InfoContext(ctx, "shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		server.Shutdown(ctx)
	}()

	slog.InfoContext(ctx, "starting server", "addr", server.Addr)
	return server.ListenAndServe()
}
