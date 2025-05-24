package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/FosteredGames/Odyssey/registry/internal/config"
	"github.com/FosteredGames/Odyssey/registry/internal/httpapi"
	"github.com/FosteredGames/Odyssey/registry/internal/registry"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/caarlos0/env/v11"
)

func loadConfig() config.Config {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		slog.Error("[registry] env parse failed", "error", err)
		os.Exit(1)
	}
	slog.Info("[registry] Config loaded", "config", cfg)
	return cfg
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	slog.Info("[registry] Starting up...")

	cfg := loadConfig()

	if cfg.PrivateKeyPath == "" {
		// Since we want to default to the user's HOME we need code to determine
		// home for various OSes.
		cfg.PrivateKeyPath = defaultKeyPath()
	}

	key, err := loadOrCreateECDSAKey(cfg.PrivateKeyPath)
	if err != nil {
		slog.Error("[registry] ecdsa key load/generation failed", "error", err)
		os.Exit(1)
	}

	slog.Info("[registry] Connecting to database...")
	db, err := data.NewDB(ctx, cfg.DBConnection)
	if err != nil {
		slog.Error("[registry] database connection failed", "error", err)
		os.Exit(1)
	}
	slog.Info("[registry] Database connection established.")

	reg := registry.NewRegistry(db, cfg.OAuth, key)

	wg := sync.WaitGroup{}
	wg.Add(1)
	var runErr error
	go func() {
		runErr = httpapi.RunRegistryServer(ctx, reg, cfg.ServerPort)
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
