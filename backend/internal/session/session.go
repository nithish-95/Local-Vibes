package session

import (
	"github.com/gorilla/sessions"
)

// Note: The key should be a long, random string for production.
var Store = sessions.NewCookieStore([]byte("super-secret-key"))
