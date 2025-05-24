package main

import (
	"context"
	"flag"
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

	var registryURL string
	flag.StringVar(&registryURL, "registry", "http://local.fosteredgames.com:8080", "Registry URL")
	flag.Parse()

	host := registry.ParseAndValidateURL(registryURL)
	fmt.Println(host)

	adminPort := GetUint16("ADMIN_PORT", 8081)
	admin := admin.New(&wg, adminPort)
	admin.Start(ctx)

	metaPort := GetUint16("META_PORT", 8082)
	meta := meta.New(&wg, metaPort)
	meta.Start(ctx)

	network := network.New()
	network.Start(ctx, &wg)

	game := game.New(&wg)
	game.Start(ctx, network.Out)

	wg.Wait()
}
