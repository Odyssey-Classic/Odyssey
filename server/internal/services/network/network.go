package network

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type ClientMap map[*websocket.Conn]*Client

type Network struct {
	clientGroup *sync.WaitGroup

	Out     chan any
	clients ClientMap
}

func New() *Network {
	return &Network{
		clientGroup: new(sync.WaitGroup),
	}
}

func (n *Network) Start(ctx context.Context, wg *sync.WaitGroup) chan any {
	wg.Add(1)
	n.clients = make(ClientMap)
	n.Out = make(chan any, 10)
	go func() {
		n.start(ctx)
		close(n.Out)
		n.clientGroup.Wait()
		wg.Done()
	}()

	return n.Out
}

func (n *Network) start(ctx context.Context) {
	server := &http.Server{
		Addr: ":8080",
	}
	server.Handler = n.wsConnect(ctx)
	server.BaseContext = func(listener net.Listener) context.Context { return ctx }

	go server.ListenAndServe()

	for {
		<-ctx.Done()
		slog.Info("network shutting down")
		server.Shutdown(ctx)
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
		client.close()
	}

	// TODO cleanup maps?
}
