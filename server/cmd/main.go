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

	"github.com/Odyssey-Classic/Odyssey/server/internal/admin"
	"github.com/Odyssey-Classic/Odyssey/server/internal/game"
	"github.com/Odyssey-Classic/Odyssey/server/internal/network"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup

	network := network.New(&wg)
	admin := admin.New(&wg)
	game := game.New(&wg)

	admin.Start(ctx)
	game.Start(ctx)
	network.Start(ctx)

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
