package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/Odyssey-Classic/Odyssey/server/internal/services/admin"
	"github.com/Odyssey-Classic/Odyssey/server/internal/services/game"
	"github.com/Odyssey-Classic/Odyssey/server/internal/services/meta"
	"github.com/Odyssey-Classic/Odyssey/server/internal/services/network"
	"github.com/Odyssey-Classic/Odyssey/server/internal/services/registry"

	"github.com/Odyssey-Classic/Odyssey/server/pb"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	var _ pb.GameMessage

	var wg sync.WaitGroup

	cfg, err := LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	host := registry.ParseAndValidateURL(cfg.RegistryURL)
	fmt.Println(host)

	admin := admin.New(&wg, uint16(cfg.AdminPort))
	admin.Start(ctx)

	meta := meta.New(&wg, uint16(cfg.MetaPort))
	meta.Start(ctx)

	network := network.New(&wg, uint16(cfg.NetworkPort))
	network.Start(ctx)

	game := game.New(&wg)
	game.Start(ctx, network.Out)

	wg.Wait()
}
