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
	"github.com/go-chi/chi/v5"
)

type Registry struct {
	DB          *data.DB
	OAuthConfig config.OAuthConfig
	PrivateKey  *ecdsa.PrivateKey
}

func (r *Registry) Run(ctx context.Context) error {
	router := chi.NewRouter()

	idServer := identity.New(r.PrivateKey, r.OAuthConfig, r.DB)
	// Mount identity endpoints at /identity/*
	router.Mount("/identity", idServer)

	serversServer := &servers.ServersServer{}
	router.Mount("/servers", idServer.AuthorizeMiddleware(serversServer))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		<-ctx.Done()

		slog.InfoContext(ctx, "shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		server.Shutdown(ctx)
	}()

	slog.InfoContext(ctx, "HTTP server starting", "address", server.Addr, "module", "registry")
	return server.ListenAndServe()
}
