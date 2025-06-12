package network

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type ClientMap map[*websocket.Conn]*Client

type Network struct {
	clientGroup *sync.WaitGroup
	wg          *sync.WaitGroup
	once        sync.Once

	Out     chan any
	clients ClientMap
	port    uint16
}

func New(wg *sync.WaitGroup, port uint16) *Network {
	return &Network{
		clientGroup: new(sync.WaitGroup),
		wg:          wg,

		clients: make(ClientMap),
		Out:     make(chan any, 10),
		port:    port,
	}
}

func (n *Network) Start(ctx context.Context) chan any {
	n.once.Do(func() {
		n.wg.Add(1)
		go func() {
			n.start(ctx)
			defer close(n.Out)
			n.clientGroup.Wait()
			n.wg.Done()
		}()
	})

	return n.Out
}

func (n *Network) start(ctx context.Context) {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", n.port),
	}
	server.Handler = n.wsConnect(ctx)
	server.BaseContext = func(listener net.Listener) context.Context { return ctx }

	// Start the server in a goroutine and capture errors
	errCh := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			slog.Error("network server error", "err", err)
		}
		errCh <- err
	}()

	for {
		<-ctx.Done()
		slog.Info("network shutting down")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			slog.Error("network shutdown error", "err", err)
		} else {
			slog.Info("network shutdown complete")
		}
		n.shutdown()
		return
	}
}

func (n *Network) addClient(ctx context.Context, client *Client) {
	slog.Info("adding client", "remote addr", client.conn.RemoteAddr())
	n.clients[client.conn] = client
	n.Out <- client
	n.processClient(ctx, client)
}

func (n *Network) shutdown() {
	slog.Info("shutting down clients")
	for _, client := range n.clients {
		err := client.close()
		if err != nil {
			slog.Error("error closing client", "error", err)
		}
	}
}
