package config

import (
	"net/url"
)

type OAuthConfig struct {
	AuthorizationURL url.URL `env:"ODY_REG_AUTHORIZATION_URL,required"`
	RedirectURL      url.URL `env:"ODY_REG_REDIRECT_URL,required"`
	RevokeURL        url.URL `env:"ODY_REG_REVOKE_URL"`
	TokenURL         url.URL `env:"ODY_REG_TOKEN_URL,required"`

	ClientID     string `env:"ODY_REG_CLIENT_ID,required"`
	ClientSecret string `env:"ODY_REG_CLIENT_SECRET,required"`
}

type Config struct {
	OAuth OAuthConfig

	PrivateKeyPath string `env:"ODY_REG_PRIVATE_KEY_PATH" envDefault:""`
	ServerPort     uint16 `env:"ODY_REG_PORT" envDefault:"8080"`
	DBConnection   string `env:"ODY_REG_DB_CONNECTION,required"`
}
