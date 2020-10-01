package hub

import (
	"github.com/dkvilo/gossip/connection"
	"github.com/dkvilo/gossip/model"
	"github.com/gorilla/websocket"
)

// Subscription structure
type Subscription struct {
	Conn *connection.Connection
	Room string
}

// Upgrader Configuration
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	Rooms map[string]map[*connection.Connection]bool

	// Inbound messages from the connections.
	Broadcast chan model.Message

	// Register requests from the connections.
	Register chan Subscription

	// Unregister requests from connections.
	Unregister chan Subscription
}

// H hub instance
var H = Hub {
	Broadcast:  make(chan model.Message),
	Register:   make(chan Subscription),
	Unregister: make(chan Subscription),
	Rooms:      make(map[string]map[*connection.Connection]bool),
}
