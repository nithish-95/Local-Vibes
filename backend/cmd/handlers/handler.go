package handlers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go Backend!"))
}

func LandingPage(w http.ResponseWriter, r *http.Request) {
	user := map[string]string{"userName": "John Doe"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
