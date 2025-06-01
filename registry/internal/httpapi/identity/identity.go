package identity

import (
	"crypto/ecdsa"
	"encoding/json"
	"net/http"

	"context"

	"github.com/FosteredGames/Odyssey/registry/internal/httpapi/identity/oauth"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity"
	"github.com/go-chi/chi/v5"
)

type API struct {
	identity *identity.Identity
	router   chi.Router
}

// UserKeyContext is the context key for the JWT user subject.
var UserKeyContext = identity.UserKeyContext

// API returns a chi.Router for all /identity endpoints, including JWKS.
func New(idServer *identity.Identity) *API {
	router := chi.NewRouter()
	oAuthServer := oauth.New(idServer.OAuthConfig(), idServer.IdentityCallback)

	router.Get("/login", oAuthServer.OAuthRedirect)
	router.Get("/oauth/callback", oAuthServer.OAuthCallback)
	router.Get("/.well-known/jwks.json", JWKSHandler(idServer.PrivateKey()))

	api := &API{
		identity: idServer,
		router:   router,
	}

	return api
}

func (a *API) Router() chi.Router {
	return a.router
}

// JWKSHandler returns a handler for /.well-known/jwks.json given an ECDSA private key.
func (a *API) JWKSHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		privateKey := a.identity.PrivateKey()
		pub := privateKey.Public().(*ecdsa.PublicKey)
		jwk := map[string]any{
			"kty": "EC",
			"crv": pub.Curve.Params().Name,
			"x":   pub.X.Text(16),
			"y":   pub.Y.Text(16),
			"use": "sig",
			"alg": "ES256",
			"kid": "1",
		}
		jwks := map[string]any{
			"keys": []any{jwk},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jwks)
	}
}

// AuthorizeMiddleware returns a middleware that validates JWTs using the provided Identity.
// If the JWT is valid, it injects the user's subject (sub) into the request context for downstream handlers.
func (a *API) AuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "no token", http.StatusUnauthorized)
			return
		}

		sub, err := a.identity.VerifyJWT(token)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// Inject the user's subject (sub) into the request context for downstream handlers
		r = r.WithContext(context.WithValue(r.Context(), UserKeyContext, sub))
		next.ServeHTTP(w, r)
	})
}
