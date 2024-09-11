package network

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func (n *Network) handler(w http.ResponseWriter, r *http.Request) {
	// Extract client metadata from JWT token
	// Use unvalidated JWT tokens for simplicity

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Info("upgrade error: %s", err)
		return
	}

	client := &Client{
		conn: c,
	}

	n.addClient(client)
	// Create new Client with conn `c`
	// Add client to clients map

	// Send client connection to Game Logic
	// Game Logic makes a "player" object with the client
	// Game Logic should ensure PC is not already playing
}
