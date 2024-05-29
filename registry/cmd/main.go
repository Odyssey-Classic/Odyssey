package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/FosteredGames/Odyssey/registry/internal/config"
	"github.com/FosteredGames/Odyssey/registry/internal/data"
	"github.com/FosteredGames/Odyssey/registry/internal/identity"
	"github.com/FosteredGames/Odyssey/registry/internal/registry"
	"golang.org/x/oauth2"
)

func main() {
	// Ignoring stop func for now as we expect this to run until killed
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	cfg, err := config.ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := data.NewDB(context.Background(), cfg.DBConnection)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	reg := &registry.Registry{
		DB:    db,
		OAuth: &cfg.OAuth,
	}

	go func() {
		reg.Run(ctx)
		wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return
		}

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

	identity := identity.New(key, oauthConf, db)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	go identity.Run(ctx)

	select {
	case <-ctx.Done():
		stop()
	}
}
