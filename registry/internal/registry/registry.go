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

	// Add HTTP request logging middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			slog.Info("[registry] HTTP request", "method", req.Method, "path", req.URL.Path, "remote", req.RemoteAddr)
			next.ServeHTTP(w, req)
		})
	})

	idServer := identity.New(r.PrivateKey, r.OAuthConfig, r.DB)
	router.Mount("/identity", idServer.Router())

	serversServer := &servers.ServersServer{}
	router.Mount("/servers", idServer.AuthorizeMiddleware(serversServer))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		<-ctx.Done()

		slog.InfoContext(ctx, "[http] shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		server.Shutdown(ctx)
	}()

	slog.InfoContext(ctx, "HTTP server starting", "address", server.Addr, "module", "registry")
	return server.ListenAndServe()
}
