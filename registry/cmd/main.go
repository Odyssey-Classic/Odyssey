package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/FosteredGames/Odyssey/registry/internal/identity"
	"golang.org/x/oauth2"
)

func main() {
	cfg, err := ConfigFromEnv()
	if err != nil {
		panic(err)
	}

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
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

	identity := identity.New(key, oauthConf)
	go identity.Run()

	for {
	}
}
