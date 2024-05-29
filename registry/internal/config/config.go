package config

import (
	"net/url"
)

const (
	authVar         = "ODY_AUTHORIZATION_URL"
	tokenVar        = "ODY_TOKEN_URL"
	revokeVar       = "ODY_REVOKE_URL"
	redirectVar     = "ODY_REDIRECT_URL"
	clientIDVar     = "ODY_CLIENT_ID"
	clientSecretVar = "ODY_CLIENT_SECRET"

	dbConnection = "ODY_DB_CONNECTION"
)

type OAuthConfig struct {
	AuthorizationURL url.URL `env:"ODY_AUTHORIZATION_URL"`
	RedirectURL      url.URL `env:"ODY_REDIRECT_URL"`
	RevokeURL        url.URL `env:"ODY_REVOKE_URL"`
	TokenURL         url.URL `env:"ODY_TOKEN_URL"`

	ClientID     string `env:"ODY_CLIENT_ID"`
	ClientSecret string `env:"ODY_CLIENT_SECRET"`
}

type Config struct {
	OAuth OAuthConfig

	DBConnection string `env:"ODY_DB_CONNECTION"`
}
