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
	"github.com/FosteredGames/Odyssey/registry/internal/registry"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/caarlos0/env/v11"
)

func main() {
	// Ignoring stop func for now as we expect this to run until killed
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	// TODO - use passed in key, provide a utility to create keys
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	db, err := data.NewDB(ctx, cfg.DBConnection)
	if err != nil {
		log.Fatal(err)
	}

	reg := &registry.Registry{
		DB:          db,
		OAuthConfig: cfg.OAuth,
		PrivateKey:  key,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
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
}
