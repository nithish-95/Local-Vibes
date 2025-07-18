package services

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/nithish-95/Local-Vibes/backend/internal/models"
)

type EventService struct {
	DB *gorm.DB
}

func NewEventService(db *gorm.DB) *EventService {
	return &EventService{DB: db}
}

func (s *EventService) CreateEvent(event *models.Event) (uint, error) {
	log.Printf("Creating event with CreatorID: %d", event.CreatorID)
	rulesJSON, err := json.Marshal(event.Rules)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal rules: %w", err)
	}
	event.Rules = string(rulesJSON)
	log.Printf("Rules string before saving: %s", event.Rules) // Added log

	result := s.DB.Create(event)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to create event: %w", result.Error)
	}

	// Automatically add the creator as a participant
	participant := struct {
		EventID uint
		UserID  uint
	}{EventID: event.ID, UserID: event.CreatorID}
	res := s.DB.Table("event_participants").Create(&participant)
	if res.Error != nil {
		// Log the error but don't fail the event creation if adding participant fails
		fmt.Printf("Warning: Failed to auto-add creator %d to event %d: %v\n", event.CreatorID, event.ID, res.Error)
	}

	return event.ID, nil
}

func (s *EventService) GetEvents() ([]models.Event, error) {
	var events []models.Event

	rows, err := s.DB.Table("events").
		Select("events.*, users.username as host_name, COUNT(event_participants.user_id) as participants").
		Joins("JOIN users ON events.creator_id = users.id").
		Joins("LEFT JOIN event_participants ON events.id = event_participants.event_id").
		Group("events.id").
		Rows()
	if err != nil {
		return nil, fmt.Errorf("failed to get events: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		var rulesJSON string
		var hostName string
		var participants int64

		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Date,
			&event.Time,
			&event.Location,
			&rulesJSON,
			&event.Capacity,
			&event.CreatorID,
			&event.ImageURL,
			&hostName,
			&participants,
		); err != nil {
			return nil, fmt.Errorf("failed to scan event: %w", err)
		}

		if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
			return nil, fmt.Errorf("failed to unmarshal rules: %w", err)
		}
		event.HostName = hostName
		event.Participants = participants
		events = append(events, event)
	}

	return events, nil
}

func (s *EventService) UpdateEvent(event *models.Event) error {
	log.Printf("Attempting to update event ID: %d with data: %+v", event.ID, event)
	rulesJSON, err := json.Marshal(event.Rules)
	if err != nil {
		return fmt.Errorf("failed to marshal rules: %w", err)
	}
	event.Rules = string(rulesJSON)

	result := s.DB.Save(event)
	if result.Error != nil {
		return fmt.Errorf("failed to update event: %w", result.Error)
	}
	log.Printf("Event ID %d updated successfully. Rows affected: %d", event.ID, result.RowsAffected)
	return nil
}

func (s *EventService) DeleteEvent(eventID uint) error {
	result := s.DB.Delete(&models.Event{}, eventID)
	if result.Error != nil {
		return fmt.Errorf("failed to delete event: %w", result.Error)
	}
	return nil
}

func (s *EventService) GetEventByID(eventID uint) (*models.Event, error) {
	var event models.Event
	var rulesJSON string
	var hostName string
	var participants int64

	row := s.DB.Table("events").
		Select("events.*, users.username as host_name, COUNT(event_participants.user_id) as participants").
		Joins("JOIN users ON events.creator_id = users.id").
		Joins("LEFT JOIN event_participants ON events.id = event_participants.event_id").
		Where("events.id = ?", eventID).
		Group("events.id").
		Row()

	err := row.Scan(
		&event.ID,
		&event.Title,
		&event.Description,
		&event.Date,
		&event.Time,
		&event.Location,
		&rulesJSON,
		&event.Capacity,
		&event.CreatorID,
		&event.ImageURL,
		&hostName,
		&participants,
	)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("event not found")
		}
		return nil, fmt.Errorf("failed to get event by ID: %w", err)
	}

	log.Printf("Rules JSON from DB for event %d: %s", eventID, rulesJSON)

	if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
		return nil, fmt.Errorf("failed to unmarshal rules for event ID %d: %w", event.ID, err)
	}
	event.HostName = hostName
	event.Participants = participants
	return &event, nil
}

func (s *EventService) GetEventsByCreatorID(creatorID uint) ([]models.Event, error) {
	log.Printf("Attempting to get events for creatorID: %d", creatorID)
	var events []models.Event

	rows, err := s.DB.Table("events").
		Select("events.*, users.username as host_name, COALESCE(COUNT(ep.user_id), 0) as participants").
		Joins("JOIN users ON events.creator_id = users.id").
		Joins("LEFT JOIN event_participants ep ON events.id = ep.event_id").
		Where("events.creator_id = ?", creatorID).
		Group("events.id").
		Rows()
	if err != nil {
		log.Printf("Error getting events by creator ID for %d: %v", creatorID, err)
		return nil, fmt.Errorf("failed to get events by creator ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		var rulesJSON string
		var hostName string
		var participants int64

		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Date,
			&event.Time,
			&event.Location,
			&rulesJSON,
			&event.Capacity,
			&event.CreatorID,
			&event.ImageURL,
			&hostName,
			&participants,
		); err != nil {
			log.Printf("Error scanning event for creator ID %d: %v", creatorID, err)
			return nil, fmt.Errorf("failed to scan event by creator ID: %w", err)
		}

		if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
			log.Printf("Error unmarshaling rules for event ID %d (creator %d): %v", event.ID, creatorID, err)
			return nil, fmt.Errorf("failed to unmarshal rules for event ID %d: %w", event.ID, err)
		}
		event.HostName = hostName
		event.Participants = participants
		events = append(events, event)
	}
	log.Printf("Found %d events for creatorID: %d", len(events), creatorID)
	return events, nil
}

func (s *EventService) GetAllEventsExcludingCreatorID(creatorID uint) ([]models.Event, error) {
	var events []models.Event

	rows, err := s.DB.Table("events").
		Select("events.*, users.username as host_name, COUNT(event_participants.user_id) as participants").
		Joins("JOIN users ON events.creator_id = users.id").
		Joins("LEFT JOIN event_participants ON events.id = event_participants.event_id").
		Where("events.creator_id != ? ", creatorID).
		Group("events.id").
		Rows()
	if err != nil {
		return nil, fmt.Errorf("failed to get all events excluding creator ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		var rulesJSON string
		var hostName string
		var participants int64

		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Date,
			&event.Time,
			&event.Location,
			&rulesJSON,
			&event.Capacity,
			&event.CreatorID,
			&event.ImageURL,
			&hostName,
			&participants,
		); err != nil {
			return nil, fmt.Errorf("failed to scan event excluding creator ID: %w", err)
		}

		if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
			return nil, fmt.Errorf("failed to unmarshal rules for event ID %d: %w", event.ID, err)
		}
		event.HostName = hostName
		event.Participants = participants
		events = append(events, event)
	}

	if events == nil {
		return []models.Event{}, nil // Return empty slice instead of nil
	}

	return events, nil
}

func (s *EventService) JoinEvent(eventID, userID uint) error {
	// Check if the user is already a participant
	var count int64
	result := s.DB.Table("event_participants").Where("event_id = ? AND user_id = ?", eventID, userID).Count(&count)
	if result.Error != nil {
		return fmt.Errorf("failed to check existing participant: %w", result.Error)
	}
	if count > 0 {
		return fmt.Errorf("user is already a participant in this event")
	}

	// Check event capacity
	var event models.Event
	result = s.DB.Select("capacity").First(&event, eventID)
	if result.Error != nil {
		return fmt.Errorf("failed to get event capacity: %w", result.Error)
	}

	var currentParticipants int64
	result = s.DB.Table("event_participants").Where("event_id = ?", eventID).Count(&currentParticipants)
	if result.Error != nil {
		return fmt.Errorf("failed to get current participants: %w", result.Error)
	}

	if event.Capacity > 0 && currentParticipants >= int64(event.Capacity) {
		return fmt.Errorf("event is full")
	}

	participant := struct {
		EventID uint
		UserID  uint
	}{EventID: eventID, UserID: userID}
	res := s.DB.Table("event_participants").Create(&participant)
	if res.Error != nil {
		return fmt.Errorf("failed to join event: %w", res.Error)
	}
	return nil
}

func (s *EventService) LeaveEvent(eventID, userID uint) error {
	result := s.DB.Table("event_participants").Where("event_id = ? AND user_id = ?", eventID, userID).Delete(nil)
	if result.Error != nil {
		return fmt.Errorf("failed to leave event: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user was not a participant in this event")
	}

	return nil
}

func (s *EventService) IsUserParticipant(eventID, userID uint) (bool, error) {
	var count int64
	result := s.DB.Table("event_participants").Where("event_id = ? AND user_id = ?", eventID, userID).Count(&count)
	if result.Error != nil {
		return false, fmt.Errorf("failed to check if user is participant: %w", result.Error)
	}
	return count > 0, nil
}

func (s *EventService) GetEventsByParticipantID(userID uint) ([]models.Event, error) {
	var events []models.Event

	rows, err := s.DB.Table("events").
		Select("events.*, users.username as host_name, COUNT(ep.user_id) as participants").
		Joins("JOIN users ON events.creator_id = users.id").
		Joins("JOIN event_participants ep_main ON events.id = ep_main.event_id").
		Joins("LEFT JOIN event_participants ep ON events.id = ep.event_id").
		Where("ep_main.user_id = ?", userID).
		Group("events.id").
		Rows()
	if err != nil {
		return nil, fmt.Errorf("failed to get events by participant ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		var rulesJSON string
		var hostName string
		var participants int64

		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Date,
			&event.Time,
			&event.Location,
			&rulesJSON,
			&event.Capacity,
			&event.CreatorID,
			&event.ImageURL,
			&hostName,
			&participants,
		); err != nil {
			return nil, fmt.Errorf("failed to scan event by participant ID: %w", err)
		}

		if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
			return nil, fmt.Errorf("failed to unmarshal rules for event ID %d: %w", event.ID, err)
		}
		event.HostName = hostName
		event.Participants = participants
		events = append(events, event)
	}

	if events == nil {
		return []models.Event{}, nil // Return empty slice instead of nil
	}

	return events, nil
}

// EventFilter defines parameters for filtering events
type EventFilter struct {
	Query       string
	StartDate   string
	EndDate     string
	CapacityMin int
	CapacityMax int
	CreatorID   uint
}

func (s *EventService) SearchEvents(filter EventFilter) ([]models.Event, error) {
	var events []models.Event

	baseQuery := s.DB.Table("events").
		Select("events.*, users.username as host_name, COUNT(event_participants.user_id) as participants").
		Joins("JOIN users ON events.creator_id = users.id").
		Joins("LEFT JOIN event_participants ON events.id = event_participants.event_id")

	if filter.Query != "" {
		baseQuery = baseQuery.Joins("JOIN events_fts ON events.id = events_fts.rowid").Where("events_fts MATCH ?", filter.Query)
	}

	if filter.StartDate != "" {
		baseQuery = baseQuery.Where("events.date >= ?", filter.StartDate)
	}
	if filter.EndDate != "" {
		baseQuery = baseQuery.Where("events.date <= ?", filter.EndDate)
	}
	if filter.CapacityMin > 0 {
		baseQuery = baseQuery.Where("events.capacity >= ?", filter.CapacityMin)
	}
	if filter.CapacityMax > 0 {
		baseQuery = baseQuery.Where("events.capacity <= ?", filter.CapacityMax)
	}
	if filter.CreatorID > 0 {
		baseQuery = baseQuery.Where("events.creator_id = ?", filter.CreatorID)
	}

	rows, err := baseQuery.Group("events.id").Rows()
	if err != nil {
		return nil, fmt.Errorf("failed to search events: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		var rulesJSON string
		var hostName string
		var participants int64

		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Date,
			&event.Time,
			&event.Location,
			&rulesJSON,
			&event.Capacity,
			&event.CreatorID,
			&event.ImageURL,
			&hostName,
			&participants,
		); err != nil {
			return nil, fmt.Errorf("failed to scan searched event: %w", err)
		}

		if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
			return nil, fmt.Errorf("failed to unmarshal rules for searched event ID %d: %w", event.ID, err)
		}
		event.HostName = hostName
		event.Participants = participants
		events = append(events, event)
	}

	return events, nil
}
