package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nithish-95/Local-Vibes/backend/internal/models"
	"github.com/nithish-95/Local-Vibes/backend/internal/session"
	"github.com/nithish-95/Local-Vibes/backend/internal/services"
)

type EventHandler struct {
	EventService *services.EventService
}

func NewEventHandler(eventService *services.EventService) *EventHandler {
	return &EventHandler{EventService: eventService}
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(int)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	// Parse multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		sendJSONError(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	var event models.Event
	event.Title = r.FormValue("title")
	event.Description = r.FormValue("description")
	event.Date = r.FormValue("date")
	event.Time = r.FormValue("time")
	event.Location = r.FormValue("location")
	event.Capacity, _ = strconv.Atoi(r.FormValue("capacity"))

	// Handle rules (JSON string from frontend)
	rulesJSON := r.FormValue("rules")
	if rulesJSON != "" {
		err = json.Unmarshal([]byte(rulesJSON), &event.Rules)
		if err != nil {
			sendJSONError(w, "Invalid rules format: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Handle image upload
	file, handler, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		uploadDir := "./uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, 0755)
		}

		filePath := filepath.Join(uploadDir, handler.Filename)
		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			sendJSONError(w, "Failed to save image: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		event.ImageURL = "/uploads/" + handler.Filename // Store relative path
	} else if err != http.ErrMissingFile {
		sendJSONError(w, "Error retrieving image file: "+err.Error(), http.StatusBadRequest)
		return
	}

	event.CreatorID = userID

	id, err := h.EventService.CreateEvent(&event)
	if err != nil {
		log.Printf("Error creating event: %v", err) // Add this line for detailed logging
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func (h *EventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.EventService.GetEvents()
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(int)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	eventID, err := strconv.Atoi(chi.URLParam(r, "eventID"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	// Parse multipart form data
	err = r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		sendJSONError(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	var event models.Event
	event.Title = r.FormValue("title")
	event.Description = r.FormValue("description")
	event.Date = r.FormValue("date")
	event.Time = r.FormValue("time")
	event.Location = r.FormValue("location")
	event.Capacity, _ = strconv.Atoi(r.FormValue("capacity"))

	// Handle rules (JSON string from frontend)
	rulesJSON := r.FormValue("rules")
	if rulesJSON != "" {
		err = json.Unmarshal([]byte(rulesJSON), &event.Rules)
		if err != nil {
			sendJSONError(w, "Invalid rules format: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Handle image upload
	file, handler, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		uploadDir := "./uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, 0755)
		}

		filePath := filepath.Join(uploadDir, handler.Filename)
		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			sendJSONError(w, "Failed to save image: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		event.ImageURL = "/uploads/" + handler.Filename // Store relative path
	} else if err != http.ErrMissingFile {
		sendJSONError(w, "Error retrieving image file: "+err.Error(), http.StatusBadRequest)
		return
	}

	event.ID = eventID

	// Verify that the user is the creator of the event
	existingEvent, err := h.EventService.GetEventByID(eventID)
	if err != nil {
		sendJSONError(w, "Event not found", http.StatusNotFound)
		return
	}
	if existingEvent.CreatorID != userID {
		sendJSONError(w, "Forbidden: You are not the creator of this event", http.StatusForbidden)
		return
	}

	if err := h.EventService.UpdateEvent(&event); err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Event updated successfully"})
}

func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(int)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	eventID, err := strconv.Atoi(chi.URLParam(r, "eventID"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	// Verify that the user is the creator of the event
	existingEvent, err := h.EventService.GetEventByID(eventID)
	if err != nil {
		sendJSONError(w, "Event not found", http.StatusNotFound)
		return	
	}
	if existingEvent.CreatorID != userID {
		sendJSONError(w, "Forbidden: You are not the creator of this event", http.StatusForbidden)
		return
	}

	if err := h.EventService.DeleteEvent(eventID); err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *EventHandler) GetHostedEvents(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(int)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	events, err := h.EventService.GetEventsByCreatorID(userID)
	if err != nil {
		log.Printf("Error getting hosted events: %v", err)
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *EventHandler) GetAvailableEvents(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(int)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	events, err := h.EventService.GetAllEventsExcludingCreatorID(userID)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(chi.URLParam(r, "eventID"))
	if err != nil {
		sendJSONError(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	event, err := h.EventService.GetEventByID(eventID)
	if err != nil {
		sendJSONError(w, "Event not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) JoinEvent(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(int)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	eventID, err := strconv.Atoi(chi.URLParam(r, "eventID"))
	if err != nil {
		sendJSONError(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	if err := h.EventService.JoinEvent(eventID, userID); err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *EventHandler) LeaveEvent(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(int)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	eventID, err := strconv.Atoi(chi.URLParam(r, "eventID"))
	if err != nil {
		sendJSONError(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	if err := h.EventService.LeaveEvent(eventID, userID); err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *EventHandler) IsJoined(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(int)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	eventID, err := strconv.Atoi(chi.URLParam(r, "eventID"))
	if err != nil {
		sendJSONError(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	isJoined, err := h.EventService.IsUserParticipant(eventID, userID)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"is_joined": isJoined})
}

func (h *EventHandler) GetJoinedEvents(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(int)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	events, err := h.EventService.GetEventsByParticipantID(userID)
	if err != nil {
		log.Printf("Error getting joined events: %v", err)
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
