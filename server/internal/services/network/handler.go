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
		if r.Header.Get("Connection") == "" || r.Header.Get("Upgrade") != "websocket" {
			http.Error(w, "WebSocket connection required", http.StatusBadRequest)
			slog.Warn("non-WebSocket connection attempt", "remote_addr", r.RemoteAddr)
			return
		}

		// Extract client metadata from JWT token
		// Use unvalidated JWT tokens for simplicity
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			slog.Warn("missing authorization token", "remote_addr", r.RemoteAddr)
			return
		}
		// Expecting format: "Bearer <token>"
		const bearerPrefix = "Bearer "
		if len(token) <= len(bearerPrefix) || token[:len(bearerPrefix)] != bearerPrefix {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			slog.Warn("invalid authorization header format", "remote_addr", r.RemoteAddr)
			return
		}
		// jwtToken := token[len(bearerPrefix):]

		// TODO: Validate JWT token here
		// If invalid, return http.Error and do not upgrade

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			slog.Error("upgrade error", "error", err)
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
