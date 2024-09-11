package network

import (
	"context"
	"log/slog"
	"sync"
)

type ClientMap map[string]*Client

type Network struct {
	wg *sync.WaitGroup

	out     chan any
	clients ClientMap
}

func New(wg *sync.WaitGroup) *Network {
	return &Network{wg: wg}
}

func (n *Network) Start(ctx context.Context) chan any {
	n.wg.Add(1)
	n.out = make(chan any, 10)
	go func() {
		n.start(ctx)
		close(n.out)
		n.wg.Done()
	}()

	return n.out
}

func (n *Network) start(ctx context.Context) {
	for {
		<-ctx.Done()
		slog.Info("network shutting down")
		return
	}
}

func (n *Network) addClient(client *Client) {
	n.clients[client.id] = client
	n.out <- client
	// n.out <- {client: id, action: "add"}
}
