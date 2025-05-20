package httpapi

import (
	"crypto/ecdsa"
	"encoding/json"
	"net/http"

	"context"

	"github.com/FosteredGames/Odyssey/registry/internal/httpapi/identity/oauth"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity"
	"github.com/go-chi/chi/v5"
)

// UserKeyContext is the context key for the JWT user subject.
var UserKeyContext = identity.UserKeyContext

// IdentityAPI returns a chi.Router for all /identity endpoints, including JWKS.
func IdentityAPI(idServer *identity.Identity) chi.Router {
	router := chi.NewRouter()
	oAuthServer := oauth.New(idServer.OAuthConfig(), idServer.IdentityCallback)

	router.Get("/login", oAuthServer.OAuthRedirect)
	router.Get("/oauth/callback", oAuthServer.OAuthCallback)
	router.Get("/.well-known/jwks.json", JWKSHandler(idServer.PrivateKey()))

	return router
}

// JWKSHandler returns a handler for /.well-known/jwks.json given an ECDSA private key.
func JWKSHandler(privateKey *ecdsa.PrivateKey) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
func AuthorizeMiddleware(identity *identity.Identity) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "no token", http.StatusUnauthorized)
				return
			}

			sub, err := identity.VerifyJWT(token)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			// Inject the user's subject (sub) into the request context for downstream handlers
			r = r.WithContext(context.WithValue(r.Context(), UserKeyContext, sub))
			handler.ServeHTTP(w, r)
		})
	}
}
