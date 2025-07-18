package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nithish-95/Local-Vibes/backend/internal/models"
	"github.com/nithish-95/Local-Vibes/backend/internal/services"
	"github.com/nithish-95/Local-Vibes/backend/internal/session"
)

type EventHandler struct {
	EventService *services.EventService
	UserService  *services.UserService
}

func NewEventHandler(eventService *services.EventService, userService *services.UserService) *EventHandler {
	return &EventHandler{EventService: eventService, UserService: userService}
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(uint)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}
	log.Printf("User ID from session: %d", userID)

	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	event.CreatorID = userID

	id, err := h.EventService.CreateEvent(&event)
	if err != nil {
		log.Printf("Error creating event: %v", err)
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]uint{"id": id})
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
	userID, ok := sess.Values["user_id"].(uint)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	eventID, err := strconv.ParseUint(chi.URLParam(r, "eventID"), 10, 64)
	if err != nil {
		sendJSONError(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	// Fetch the existing event to preserve CreatorID
	existingEvent, err := h.EventService.GetEventByID(uint(eventID))
	if err != nil {
		sendJSONError(w, "Event not found", http.StatusNotFound)
		return
	}

	// Verify that the user is the creator of the event
	if existingEvent.CreatorID != userID {
		sendJSONError(w, "Forbidden: You are not the creator of this event", http.StatusForbidden)
		return
	}

	var updatedEventData models.Event
	err = json.NewDecoder(r.Body).Decode(&updatedEventData)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update fields from the request body, but preserve the original CreatorID
	existingEvent.Title = updatedEventData.Title
	existingEvent.Description = updatedEventData.Description
	existingEvent.Date = updatedEventData.Date
	existingEvent.Time = updatedEventData.Time
	existingEvent.Location = updatedEventData.Location
	existingEvent.Rules = updatedEventData.Rules
	existingEvent.Capacity = updatedEventData.Capacity
	existingEvent.ImageURL = updatedEventData.ImageURL

	log.Printf("Handler: Received update request for event ID: %d, Data: %+v", eventID, existingEvent) // Added log

	if err := h.EventService.UpdateEvent(existingEvent); err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Event updated successfully"})
}

func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(uint)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	eventID, err := strconv.ParseUint(chi.URLParam(r, "eventID"), 10, 64)
	if err != nil {
		sendJSONError(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	// Verify that the user is the creator of the event
	existingEvent, err := h.EventService.GetEventByID(uint(eventID))
	if err != nil {
		sendJSONError(w, "Event not found", http.StatusNotFound)
		return
	}
	if existingEvent.CreatorID != userID {
		sendJSONError(w, "Forbidden: You are not the creator of this event", http.StatusForbidden)
		return
	}

	if err := h.EventService.DeleteEvent(uint(eventID)); err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *EventHandler) GetHostedEvents(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(uint)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		log.Printf("Error getting user for hosted events: %v", err)
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("GetHostedEvents: User ID from session: %d, Username: %s", userID, user.Username)

	events, err := h.EventService.GetEventsByCreatorID(userID)
	if err != nil {
		log.Printf("Error getting hosted events: %v", err)
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if events == nil {
		json.NewEncoder(w).Encode([]models.Event{}) // Encode empty array if nil
	} else {
		json.NewEncoder(w).Encode(events)
	}
}

func (h *EventHandler) GetAvailableEvents(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(uint)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	events, err := h.EventService.GetAllEventsExcludingCreatorID(userID)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Events for GetAvailableEvents: %+v", events)

	w.Header().Set("Content-Type", "application/json")
	if events == nil {
		json.NewEncoder(w).Encode([]models.Event{}) // Ensure empty array is returned
	} else {
		json.NewEncoder(w).Encode(events)
	}
}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.ParseUint(chi.URLParam(r, "eventID"), 10, 64)
	if err != nil {
		sendJSONError(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	event, err := h.EventService.GetEventByID(uint(eventID))
	if err != nil {
		sendJSONError(w, "Event not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) JoinEvent(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(uint)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	eventID, err := strconv.ParseUint(chi.URLParam(r, "eventID"), 10, 64)
	if err != nil {
		sendJSONError(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	if err := h.EventService.JoinEvent(uint(eventID), userID); err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *EventHandler) LeaveEvent(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(uint)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	eventID, err := strconv.ParseUint(chi.URLParam(r, "eventID"), 10, 64)
	if err != nil {
		sendJSONError(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	if err := h.EventService.LeaveEvent(uint(eventID), userID); err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *EventHandler) IsJoined(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(uint)
	if !ok {
		sendJSONError(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	eventID, err := strconv.ParseUint(chi.URLParam(r, "eventID"), 10, 64)
	if err != nil {
		sendJSONError(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	isJoined, err := h.EventService.IsUserParticipant(uint(eventID), userID)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"is_joined": isJoined})
}

func (h *EventHandler) GetJoinedEvents(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.Store.Get(r, "session-name")
	userID, ok := sess.Values["user_id"].(uint)
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

	log.Printf("Events for GetJoinedEvents: %+v", events)

	w.Header().Set("Content-Type", "application/json")
	if events == nil {
		json.NewEncoder(w).Encode([]models.Event{}) // Ensure empty array is returned
	} else {
		json.NewEncoder(w).Encode(events)
	}
}

func (h *EventHandler) SearchEvents(w http.ResponseWriter, r *http.Request) {
	var filter services.EventFilter

	// Parse query parameters
	filter.Query = r.URL.Query().Get("q")
	filter.StartDate = r.URL.Query().Get("startDate")
	filter.EndDate = r.URL.Query().Get("endDate")

	if capacityMinStr := r.URL.Query().Get("capacityMin"); capacityMinStr != "" {
		val, err := strconv.Atoi(capacityMinStr)
		if err == nil {
			filter.CapacityMin = val
		}
	}
	if capacityMaxStr := r.URL.Query().Get("capacityMax"); capacityMaxStr != "" {
		val, err := strconv.Atoi(capacityMaxStr)
		if err == nil {
			filter.CapacityMax = val
		}
	}

	// Optional: If you want to filter by creator_id from query params
	if creatorIDStr := r.URL.Query().Get("creatorId"); creatorIDStr != "" {
		val, err := strconv.ParseUint(creatorIDStr, 10, 64)
		if err == nil {
			filter.CreatorID = uint(val)
		}
	}

	events, err := h.EventService.SearchEvents(filter)
	if err != nil {
		log.Printf("Error searching events: %v", err)
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
