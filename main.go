package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dkvilo/gossip/controller"
	"github.com/dkvilo/gossip/functions"
	"github.com/dkvilo/gossip/hub"
	"github.com/dkvilo/gossip/middleware"
)

func init() {
	go hub.H.Run()
}

func main()  {

	ctr := controller.New()
	http.HandleFunc("/", ctr.StaticServer)
	http.HandleFunc("/ws", middleware.VerifyHmac(ctr.WebSocket))
	
	fmt.Println("Hmac accessToken:", functions.GenerateHmac(os.Getenv("HMAC_MESSAGE"), os.Getenv("HMAC_SECRET")))
	panic(http.ListenAndServe(":3000", nil))
}
