package httpapi

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/FosteredGames/Odyssey/registry/internal/httpapi/identity"
	"github.com/FosteredGames/Odyssey/registry/internal/httpapi/servers"
	"github.com/FosteredGames/Odyssey/registry/internal/registry"
	"github.com/go-chi/chi/v5"
)

// loggingMiddleware logs HTTP requests.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		slog.Info("[registry] HTTP request", "method", req.Method, "path", req.URL.Path, "remote", req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}

// RunRegistryServer sets up and runs the Odyssey registry HTTP server using the Registry for business logic.
func RunRegistryServer(ctx context.Context, reg *registry.Registry, port uint16) error {
	router := chi.NewRouter()

	router.Use(loggingMiddleware)

	idAPI := identity.New(reg.IdentityService())
	router.Get("/.well-known/jwks.json", idAPI.JWKSHandler())
	router.Mount("/identity", idAPI.Router())

	serversAPI := servers.New(reg.ServersService())
	router.Mount("/servers", idAPI.AuthorizeMiddleware(serversAPI.Router()))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	go func() {
		<-ctx.Done()

		slog.InfoContext(ctx, "[http] shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		server.Shutdown(ctx)
	}()

	slog.InfoContext(ctx, "[http] HTTP server starting", "address", server.Addr, "module", "registry")

	return server.ListenAndServe()
}
