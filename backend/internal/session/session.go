package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Note: The key should be a long, random string for production.
var Store = sessions.NewCookieStore([]byte("super-secret-key"))

func init() {
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   false, // Set to true in production with HTTPS
		Domain:   "",    // Allow cookie to be sent to different subdomains/IPs for development
	}
}
