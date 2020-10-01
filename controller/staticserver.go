package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// StaticServer Controller
func (c *Controller) StaticServer(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.ServeFile(w, r, fmt.Sprintf("%s/public/index.html", dir));
}
