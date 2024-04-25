package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"os"
	"os/signal"

	"github.com/FosteredGames/Odyssey/registry/internal/identity"
	"golang.org/x/oauth2"
)

func main() {
	cfg, err := ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	go identity.Run(ctx)

	select {
	case <-ctx.Done():
		stop()
	}
}
