package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nithish-95/Local-Vibes/backend/internal/models"
	"github.com/nithish-95/Local-Vibes/backend/internal/services"
	"github.com/nithish-95/Local-Vibes/backend/internal/session"
)

type AuthHandler struct {
	UserService *services.UserService
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{UserService: userService}
}

// sendJSONError sends a JSON error response.
func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.UserService.RegisterUser(user.Username, user.Password); err != nil {
		log.Printf("Error registering user: %v", err)
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := h.UserService.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	sess, _ := session.Store.Get(r, "session-name")
	sess.Values["user_id"] = userID
	sess.Save(r, w)

	w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	sess.Options.MaxAge = -1
	sess.Save(r, w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}

func (h *AuthHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(int)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
