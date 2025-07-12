package main

import (
	"log"
	"net/http"

	"github.com/nithish-95/Local-Vibes/backend/internal/database"
	"github.com/nithish-95/Local-Vibes/backend/internal/routes"
)

func main() {
	database.InitDB()
	r := routes.SetupRouter()

	// Serve static files from the "uploads" directory
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
