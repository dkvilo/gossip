package hub

import "github.com/dkvilo/gossip/connection"

// Run - Starts the hub
func (h *Hub) Run() {
  for {
		select {
      case s := <-h.Register:
				connections := h.Rooms[s.Room]
				if connections == nil {
					connections = make(map[*connection.Connection]bool)
					h.Rooms[s.Room] = connections
				}
				h.Rooms[s.Room][s.Conn] = true

      case s := <-h.Unregister:
				connections := h.Rooms[s.Room]
				if connections != nil {
					if _, ok := connections[s.Conn]; ok {
						delete(connections, s.Conn)
						close(s.Conn.Send)
						
						if len(connections) == 0 {
							delete(h.Rooms, s.Room)
						}
					}
				}
      case m := <-h.Broadcast:
        connections := h.Rooms[m.Room]
        for c := range connections {
          select {
						case c.Send <- m.Data:
						default:
							close(c.Send)
							delete(connections, c)
							
							if len(connections) == 0 {
								delete(h.Rooms, m.Room)
							}
          }
        }
		}
  }
}
