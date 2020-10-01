package connection

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	// WriteWait - Time allowed to write a message to the peer.
	WriteWait = 10 * time.Second

	// PongWait - Time allowed to read the next pong message from the peer.
	PongWait = 60 * time.Second

	// PingPeriod - Send pings to peer with this period. Must be less than pongWait.
	PingPeriod = (PongWait * 9) / 10

	// MaxMessageSize - Maximum message size allowed from peer.
	MaxMessageSize = 1024
)

// Connection is an middleman between the websocket connection and the hub.
type Connection struct {
	// Ws - The websocket connection.
	Ws *websocket.Conn

	// Send - Buffered channel of outbound messages.
	Send chan []byte
}

// Write - writes a message with the given message type and payload.
func (c *Connection) Write(mt int, payload []byte) error {
	c.Ws.SetWriteDeadline(time.Now().Add(WriteWait))
	return c.Ws.WriteMessage(mt, payload)
}