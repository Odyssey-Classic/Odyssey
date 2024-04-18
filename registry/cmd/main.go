package main

import (
	"github.com/FosteredGames/Odyssey/registry/internal/oauth"
	"golang.org/x/oauth2"
)

func main() {
	cfg, err := ConfigFromEnv()
	if err != nil {
		panic(err)
	}

	oauthConf := &oauth2.Config{
		ClientID:     cfg.OAuth.ClientID,
		ClientSecret: cfg.OAuth.ClientSecret,
		Scopes:       []string{},
		RedirectURL:  cfg.OAuth.RedirectURL.String(),
		Endpoint: oauth2.Endpoint{
			AuthURL:  cfg.OAuth.AuthorizationURL.String(),
			TokenURL: cfg.OAuth.TokenURL.String(),
		},
	}

	oauth.Run(oauthConf)
}
