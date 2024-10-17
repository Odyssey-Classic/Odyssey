package network

import (
	"context"
	"log/slog"

	"github.com/Odyssey-Classic/Odyssey/server/pb"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

// Represents a client with a WebSocket connection
type Client struct {
	conn       *websocket.Conn
	fromRemote chan any
	toRemote   chan any

	closed bool
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:       conn,
		fromRemote: make(chan any, 10),
		toRemote:   make(chan any, 10),
	}
}

func (c *Client) close() error {
	return c.conn.Close()
}

// Reads a single message
func (c *Client) read() (any, error) {
	_, bytes, err := c.conn.ReadMessage()
	if err != nil {
		slog.Error(err.Error())
	}

	msg := &pb.GameMessage{}
	err = proto.Unmarshal(bytes, msg)
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("received message", "type", msg.Type)

	return msg, err
}

// Writes a single message
func (c *Client) write(msg any) error {
	slog.Debug("writing message", "message", msg)
	return nil
}

// Infinite loop that sends messages to remote
func (c *Client) processOutbound(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			c.close()
			return nil
		case msg := <-c.toRemote:
			err := c.write(msg)
			if err != nil {
				c.close()
				return err
			}
		}
	}
}

// Infinite loop that receives messages from remote
func (c *Client) processInbound(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			c.close()
			return nil
		default:
			msg, err := c.read()
			if err != nil {
				c.close()
				return err
			}
			c.fromRemote <- msg
		}
	}
}
