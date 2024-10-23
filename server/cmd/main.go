package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"

	// "github.com/Odyssey-Classic/Odyssey/server/internal/services/admin"
	"github.com/Odyssey-Classic/Odyssey/server/internal/services/game"
	"github.com/Odyssey-Classic/Odyssey/server/internal/services/network"

	"github.com/Odyssey-Classic/Odyssey/server/pb"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var _ pb.GameMessage

	var wg sync.WaitGroup

	network := network.New()
	// admin := admin.New(&wg)
	game := game.New(&wg)

	// admin.Start(ctx)
	network.Start(ctx, &wg)
	game.Start(ctx, network.Out)

	var registry string
	flag.StringVar(&registry, "registry", "http://local.fosteredgames.com:8080", "Registry URL")
	flag.Parse()

	host, err := url.Parse(registry)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(host)

	wg.Wait()
}
