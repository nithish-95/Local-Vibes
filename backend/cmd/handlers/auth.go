package handlers

import (
	"encoding/json"
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

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.UserService.RegisterUser(user.Username, user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := h.UserService.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	sess, _ := session.Store.Get(r, "session-name")
	sess.Values["user_id"] = userID
	sess.Save(r, w)

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	sess.Options.MaxAge = -1
	sess.Save(r, w)
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(int)
	if !ok {
		http.Error(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"userName": user.Username})
}
