package controller

import (
	"net/http"

	"github.com/dkvilo/gossip/client"
	"github.com/dkvilo/gossip/hub"
)

// WebSocket Controller
func (c *Controller) WebSocket(w http.ResponseWriter, r *http.Request)  {
	client.Serve(&hub.H, w, r);
}
