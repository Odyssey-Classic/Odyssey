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

	network := network.New()
	adminPort := GetUint16("ADMIN_PORT", 8081)
	metaPort := GetUint16("META_PORT", 8082)

	admin := admin.New(&wg, adminPort)
	meta := meta.New(&wg, metaPort)
	game := game.New(&wg)

	admin.Start(ctx)
	meta.Start(ctx)
	network.Start(ctx, &wg)
	game.Start(ctx, network.Out)

	var registryURL string
	flag.StringVar(&registryURL, "registry", "http://local.fosteredgames.com:8080", "Registry URL")
	flag.Parse()

	host := registry.ParseAndValidateURL(registryURL)

	fmt.Println(host)

	wg.Wait()
}
