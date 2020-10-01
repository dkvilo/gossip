package hub

import (
	"log"
	"time"

	"github.com/dkvilo/gossip/connection"
	"github.com/dkvilo/gossip/model"
	"github.com/gorilla/websocket"
)

// ReadPump pumps messages from the websocket connection to the hub.
// reads from this goroutine.
func (s Subscription) ReadPump() {
	c := s.Conn

	defer func() {	
		H.Unregister <- s
		c.Ws.Close()
	}()
	
	c.Ws.SetReadLimit(connection.MaxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(connection.PongWait))
	c.Ws.SetPongHandler(func(string) error { c.Ws.SetReadDeadline(time.Now().Add(connection.PongWait)); return nil })
	
	for {
		_, msg, err := c.Ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		
		H.Broadcast <- model.Message{Data: msg, Room: s.Room}
	}
}