package network

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func init() {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
}

func (n *Network) wsConnect(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract client metadata from JWT token
		// Use unvalidated JWT tokens for simplicity

		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			slog.Info("upgrade error:", "error", err)
			return
		}

		client := &Client{
			conn: conn,
		}

		n.addClient(ctx, client)
		// Create new Client with conn `c`
		// Add client to clients map

		// Send client connection to Game Logic
		// Game Logic makes a "player" object with the client
		// Game Logic should ensure PC is not already playing
	}
}
