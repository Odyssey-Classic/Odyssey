package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/FosteredGames/Odyssey/registry/internal/config"
	"github.com/FosteredGames/Odyssey/registry/internal/registry"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/caarlos0/env/v11"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	slog.Info("[registry] Starting up...")

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		slog.Error("[registry] env parse failed", "error", err)
		os.Exit(1)
	}

	slog.Info("[registry] Connecting to database...")
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		slog.Error("[registry] ecdsa key generation failed", "error", err)
		os.Exit(1)
	}

	db, err := data.NewDB(ctx, cfg.DBConnection)
	if err != nil {
		slog.Error("[registry] database connection failed", "error", err)
		os.Exit(1)
	}
	slog.Info("[registry] Database connection established.")

	reg := &registry.Registry{
		DB:          db,
		OAuthConfig: cfg.OAuth,
		PrivateKey:  key,
	}

	slog.Info("[registry] Server starting on :8080")

	wg := sync.WaitGroup{}
	wg.Add(1)
	var runErr error
	go func() {
		runErr = reg.Run(ctx)
		wg.Done()
	}()

	<-ctx.Done()
	slog.Info("[registry] Shutting down...")
	wg.Wait()
	if runErr != nil && !errors.Is(runErr, http.ErrServerClosed) {
		slog.Error("[registry] exited with error", "error", runErr)
		os.Exit(1)
	}
	os.Exit(0)
}
