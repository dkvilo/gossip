package middleware

import (
	"net/http"
	"os"

	"github.com/dkvilo/gossip/functions"
)

// VerifyHmac - checks if client is authenticated
func VerifyHmac(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/json")
		if r.URL.Query().Get("accessToken") != "" {
			if ok := functions.ValidMAC(os.Getenv("HMAC_MESSAGE"), r.URL.Query().Get("accessToken"), os.Getenv("HMAC_SECRET")); ok {
				next(w, r)
			} else {
				http.Error(w, "Authenticated failed", http.StatusNetworkAuthenticationRequired)
				return
			}
		} else {
			http.Error(w, "accessToken is missing", http.StatusNetworkAuthenticationRequired)
			return
		}
	}
}
