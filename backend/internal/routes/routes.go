package routes

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors" // Import the cors package
	"github.com/nithish-95/Local-Vibes/backend/cmd/handlers"
	"github.com/nithish-95/Local-Vibes/backend/internal/database"
	"github.com/nithish-95/Local-Vibes/backend/internal/services"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Configure CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins for development
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Initialize services
	userService := services.NewUserService(database.DB)
	eventService := services.NewEventService(database.DB)

	// Initialize handlers with services
	authHandler := handlers.NewAuthHandler(userService)
	eventHandler := handlers.NewEventHandler(eventService, userService)

	// Public routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go Backend!"))
	})

	// Auth routes
	r.Post("/api/register", authHandler.Register)
	r.Post("/api/login", authHandler.Login)
	r.Post("/api/logout", authHandler.Logout)
	r.Get("/api/user", authHandler.GetCurrentUser)

	// Event routes
	r.Post("/api/events", eventHandler.CreateEvent)
	r.Get("/api/events", eventHandler.GetEvents)
	r.Get("/api/events/{eventID}", eventHandler.GetEventByID)
	r.Put("/api/events/{eventID}", eventHandler.UpdateEvent)
	r.Delete("/api/events/{eventID}", eventHandler.DeleteEvent)
	r.Get("/api/events/hosted", eventHandler.GetHostedEvents)
	r.Get("/api/events/available", eventHandler.GetAvailableEvents)
	r.Post("/api/events/{eventID}/join", eventHandler.JoinEvent)
	r.Post("/api/events/{eventID}/leave", eventHandler.LeaveEvent)
	r.Get("/api/events/{eventID}/is-joined", eventHandler.IsJoined)
	r.Get("/api/events/joined", eventHandler.GetJoinedEvents)
	r.Get("/api/events/search", eventHandler.SearchEvents)

	return r
}

