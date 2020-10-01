package client

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dkvilo/gossip/connection"
	"github.com/dkvilo/gossip/hub"
)

// Serve handles websocket requests from the peer.
func Serve(h *hub.Hub, w http.ResponseWriter, r *http.Request) {

	var room string = r.URL.Query().Get("room")
	hub.Upgrader.CheckOrigin = func(r *http.Request) bool {
		originList := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
		for originIndex := range originList {
			if r.Header.Get("Origin") == originList[originIndex] {
				return true
			}
		}
		return false
	}

	conn, err := hub.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	sub := hub.Subscription{
		Conn: &connection.Connection{Send: make(chan []byte, 256), Ws: conn},
		Room: room,
	}
	
	hub.H.Register <- sub
 
	go sub.WritePump()
	go sub.ReadPump()
}

