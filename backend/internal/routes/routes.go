package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/nithish-95/Local-Vibes/backend/cmd/handlers"
	"github.com/nithish-95/Local-Vibes/backend/internal/database"
	"github.com/nithish-95/Local-Vibes/backend/internal/session"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Add routes here later
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go Backend!"))
	})

	// Corrected route
	r.Get("/api/user", func(w http.ResponseWriter, r *http.Request) {
		sess, _ := session.Store.Get(r, "session-name")
		userID, ok := sess.Values["user_id"].(int)
		if !ok {
			http.Error(w, "Not authenticated", http.StatusUnauthorized)
			return
		}

		var username string
		err := database.DB.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
		if err != nil {
			http.Error(w, "User not found", http.StatusInternalServerError)
			return
		}

		user := map[string]string{"userName": username}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	r.Post("/api/register", handlers.Register)
	r.Post("/api/login", handlers.Login)
	r.Post("/api/logout", handlers.Logout)

	return r
}

