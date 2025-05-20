package httpapi

import (
	"crypto/ecdsa"
	"encoding/json"
	"net/http"

	"github.com/FosteredGames/Odyssey/registry/internal/httpapi/identity/oauth"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity"
	"github.com/go-chi/chi/v5"
)

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
