package session

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

// Note: The key should be a long, random string for production.
var Store *sessions.CookieStore

func init() {
	secretKey := os.Getenv("SESSION_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SESSION_SECRET_KEY environment variable not set")
	}
	Store = sessions.NewCookieStore([]byte(secretKey))

	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false, // Set to true in production with HTTPS
		Domain:   "",    // Allow cookie to be sent to different subdomains/IPs for development
	}
}


