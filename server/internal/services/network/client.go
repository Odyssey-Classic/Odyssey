package network

import "github.com/gorilla/websocket"

// Represents a client with a WebSocket connection
type Client struct {
	id   string
	conn *websocket.Conn
	in   chan any
	out  chan any
}
