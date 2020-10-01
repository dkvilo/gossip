package functions

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// ValidMAC reports whether messageMAC is a valid HMAC tag for message.
func ValidMAC(message, messageMAC, key string) bool {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	byteMAC, _ := hex.DecodeString(messageMAC)
	return hmac.Equal(byteMAC, expectedMAC)
}

// GenerateHmac - Generates Hmac hex string
func GenerateHmac(message, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
