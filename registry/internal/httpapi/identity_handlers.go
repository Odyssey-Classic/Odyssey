package httpapi

import (
	"github.com/FosteredGames/Odyssey/registry/internal/httpapi/identity/oauth"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity"
	"github.com/go-chi/chi/v5"
)

// NewIdentityRouter sets up the HTTP routes for identity operations and returns a chi.Router.
func NewIdentityRouter(idServer *identity.IdentityServer) chi.Router {
	router := chi.NewRouter()
	oAuthServer := oauth.New(idServer.OAuthConfig(), idServer.IdentityCallback)

	router.Get("/login", oAuthServer.OAuthRedirect)
	router.Get("/oauth/callback", oAuthServer.OAuthCallback)
	router.Get("/.well-known/jwks.json", JWKSHandler(idServer.PrivateKey()))

	return router
}
