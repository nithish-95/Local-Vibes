package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nithish-95/Local-Vibes/backend/internal/models"
	"log"
	"strings"
)

type EventService struct {
	DB *sql.DB
}

func NewEventService(db *sql.DB) *EventService {
	return &EventService{DB: db}
}

func (s *EventService) CreateEvent(event *models.Event) (int64, error) {
	rulesJSON, err := json.Marshal(event.Rules)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal rules: %w", err)
	}

	query := `INSERT INTO events (title, description, date, time, location, rules, capacity, image_url, creator_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := s.DB.Exec(query,
		event.Title,
		event.Description,
		event.Date,
		event.Time,
		event.Location,
		string(rulesJSON),
		event.Capacity,
		event.ImageURL,
		event.CreatorID,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create event: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	// Automatically add the creator as a participant
	_, err = s.DB.Exec("INSERT INTO event_participants (event_id, user_id) VALUES (?, ?)", id, event.CreatorID)
	if err != nil {
		// Log the error but don't fail the event creation if adding participant fails
		fmt.Printf("Warning: Failed to auto-add creator %d to event %d: %v\n", event.CreatorID, id, err)
	}

	return id, nil
}

func (s *EventService) GetEvents() ([]models.Event, error) {
	rows, err := s.DB.Query("SELECT e.id, e.title, e.description, e.date, e.time, e.location, e.rules, e.capacity, e.creator_id, e.image_url, u.username, COALESCE(COUNT(ep.user_id), 0) FROM events e JOIN users u ON e.creator_id = u.id LEFT JOIN event_participants ep ON e.id = ep.event_id GROUP BY e.id")
	if err != nil {
		return nil, fmt.Errorf("failed to get events: %w", err)
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		var rulesJSON string
		var hostName string
		if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Time, &event.Location, &rulesJSON, &event.Capacity, &event.CreatorID, &event.ImageURL, &hostName, &event.Participants); err != nil {
			return nil, fmt.Errorf("failed to scan event: %w", err)
		}

		if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
			return nil, fmt.Errorf("failed to unmarshal rules: %w", err)
		}
		event.HostName = hostName
		events = append(events, event)
	}

	return events, nil
}

func (s *EventService) UpdateEvent(event *models.Event) error {
	rulesJSON, err := json.Marshal(event.Rules)
	if err != nil {
		return fmt.Errorf("failed to marshal rules: %w", err)
	}

	query := `UPDATE events SET title = ?, description = ?, date = ?, time = ?, location = ?, rules = ?, capacity = ?, image_url = ? WHERE id = ?`
	_, err = s.DB.Exec(query,
		event.Title,
		event.Description,
		event.Date,
		event.Time,
		event.Location,
		string(rulesJSON),
		event.Capacity,
		event.ImageURL,
		event.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update event: %w", err)
	}
	return nil
}

func (s *EventService) DeleteEvent(eventID int) error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := s.DB.Exec(query, eventID)
	if err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}
	return nil
}

func (s *EventService) GetEventByID(eventID int) (*models.Event, error) {
	var event models.Event
	var rulesJSON string
	var hostName string
	query := `SELECT e.id, e.title, e.description, e.date, e.time, e.location, e.rules, e.capacity, e.creator_id, e.image_url, u.username, COALESCE(COUNT(ep.user_id), 0) FROM events e JOIN users u ON e.creator_id = u.id LEFT JOIN event_participants ep ON e.id = ep.event_id WHERE e.id = ? GROUP BY e.id`
	err := s.DB.QueryRow(query, eventID).Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Time, &event.Location, &rulesJSON, &event.Capacity, &event.CreatorID, &event.ImageURL, &hostName, &event.Participants)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("event not found")
		}
		return nil, fmt.Errorf("failed to get event by ID: %w", err)
	}

	log.Printf("Rules JSON from DB for event %d: %s", eventID, rulesJSON) // Added log

	if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
		return nil, fmt.Errorf("failed to unmarshal rules for event ID %d: %w", event.ID, err)
	}
	event.HostName = hostName
	return &event, nil
}

func (s *EventService) GetEventsByCreatorID(creatorID int) ([]models.Event, error) {
	rows, err := s.DB.Query("SELECT e.id, e.title, e.description, e.date, e.time, e.location, e.rules, e.capacity, e.creator_id, e.image_url, u.username, COALESCE(COUNT(ep.user_id), 0) FROM events e JOIN users u ON e.creator_id = u.id LEFT JOIN event_participants ep ON e.id = ep.event_id WHERE e.creator_id = ? GROUP BY e.id", creatorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get events by creator ID: %w", err)
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		var rulesJSON string
		var hostName string
		if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Time, &event.Location, &rulesJSON, &event.Capacity, &event.CreatorID, &event.ImageURL, &hostName, &event.Participants); err != nil {
			return nil, fmt.Errorf("failed to scan event by creator ID: %w", err)
		}

		if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
			return nil, fmt.Errorf("failed to unmarshal rules for event ID %d: %w", event.ID, err)
		}
		event.HostName = hostName
		events = append(events, event)
	}

	return events, nil
}

func (s *EventService) GetAllEventsExcludingCreatorID(creatorID int) ([]models.Event, error) {
	rows, err := s.DB.Query("SELECT e.id, e.title, e.description, e.date, e.time, e.location, e.rules, e.capacity, e.creator_id, e.image_url, u.username, COALESCE(COUNT(ep.user_id), 0) FROM events e JOIN users u ON e.creator_id = u.id LEFT JOIN event_participants ep ON e.id = ep.event_id WHERE e.creator_id != ? GROUP BY e.id", creatorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all events excluding creator ID: %w", err)
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		var rulesJSON string
		var hostName string
		if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Time, &event.Location, &rulesJSON, &event.Capacity, &event.CreatorID, &event.ImageURL, &hostName, &event.Participants); err != nil {
			return nil, fmt.Errorf("failed to scan event excluding creator ID: %w", err)
		}

		if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
			return nil, fmt.Errorf("failed to unmarshal rules for event ID %d: %w", event.ID, err)
		}
		event.HostName = hostName
		events = append(events, event)
	}

	return events, nil
}

func (s *EventService) JoinEvent(eventID, userID int) error {
	// Check if the user is already a participant
	var count int
	err := s.DB.QueryRow("SELECT COUNT(*) FROM event_participants WHERE event_id = ? AND user_id = ?", eventID, userID).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check existing participant: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("user is already a participant in this event")
	}

	// Check event capacity
	var capacity int
	var currentParticipants int
	err = s.DB.QueryRow("SELECT capacity, COUNT(ep.user_id) FROM events e LEFT JOIN event_participants ep ON e.id = ep.event_id WHERE e.id = ? GROUP BY e.id", eventID).Scan(&capacity, &currentParticipants)
	if err != nil {
		return fmt.Errorf("failed to get event capacity: %w", err)
	}

	if capacity > 0 && currentParticipants >= capacity {
		return fmt.Errorf("event is full")
	}

	query := `INSERT INTO event_participants (event_id, user_id) VALUES (?, ?)`
	_, err = s.DB.Exec(query, eventID, userID)
	if err != nil {
		return fmt.Errorf("failed to join event: %w", err)
	}
	return nil
}

func (s *EventService) LeaveEvent(eventID, userID int) error {
	query := `DELETE FROM event_participants WHERE event_id = ? AND user_id = ?`
	result, err := s.DB.Exec(query, eventID, userID)
	if err != nil {
		return fmt.Errorf("failed to leave event: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user was not a participant in this event")
	}

	return nil
}

func (s *EventService) IsUserParticipant(eventID, userID int) (bool, error) {
	var count int
	err := s.DB.QueryRow("SELECT COUNT(*) FROM event_participants WHERE event_id = ? AND user_id = ?", eventID, userID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check if user is participant: %w", err)
	}
	return count > 0, nil
}

func (s *EventService) GetEventsByParticipantID(userID int) ([]models.Event, error) {
	rows, err := s.DB.Query(`
		SELECT
			e.id, e.title, e.description, e.date, e.time, e.location, e.rules, e.capacity, e.creator_id, e.image_url, u.username, COALESCE(COUNT(ep.user_id), 0)
		FROM events e
		JOIN users u ON e.creator_id = u.id
		JOIN event_participants ep_main ON e.id = ep_main.event_id
		LEFT JOIN event_participants ep ON e.id = ep.event_id
		WHERE ep_main.user_id = ?
		GROUP BY e.id
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get events by participant ID: %w", err)
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		var rulesJSON string
		var hostName string
		if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Time, &event.Location, &rulesJSON, &event.Capacity, &event.CreatorID, &event.ImageURL, &hostName, &event.Participants); err != nil {
			return nil, fmt.Errorf("failed to scan event by participant ID: %w", err)
		}

		if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
			return nil, fmt.Errorf("failed to unmarshal rules for event ID %d: %w", event.ID, err)
		}
		event.HostName = hostName
		events = append(events, event)
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
	CreatorID   int
}

func (s *EventService) SearchEvents(filter EventFilter) ([]models.Event, error) {
	var args []interface{}
	query := `
		SELECT
			e.id, e.title, e.description, e.date, e.time, e.location, e.rules, e.capacity, e.creator_id, e.image_url, u.username, COALESCE(COUNT(ep.user_id), 0)
		FROM events e
		JOIN users u ON e.creator_id = u.id
		LEFT JOIN event_participants ep ON e.id = ep.event_id
	`

	whereClauses := []string{}

	if filter.Query != "" {
		query += ` JOIN events_fts ON e.id = events_fts.rowid`
		whereClauses = append(whereClauses, "events_fts MATCH ?")
		args = append(args, filter.Query)
	}

	if filter.StartDate != "" {
		whereClauses = append(whereClauses, "e.date >= ?")
		args = append(args, filter.StartDate)
	}
	if filter.EndDate != "" {
		whereClauses = append(whereClauses, "e.date <= ?")
		args = append(args, filter.EndDate)
	}
	if filter.CapacityMin > 0 {
		whereClauses = append(whereClauses, "e.capacity >= ?")
		args = append(args, filter.CapacityMin)
	}
	if filter.CapacityMax > 0 {
		whereClauses = append(whereClauses, "e.capacity <= ?")
		args = append(args, filter.CapacityMax)
	}
	if filter.CreatorID > 0 {
		whereClauses = append(whereClauses, "e.creator_id = ?")
		args = append(args, filter.CreatorID)
	}

	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query += ` GROUP BY e.id`

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search events: %w", err)
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		var rulesJSON string
		var hostName string
		if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Time, &event.Location, &rulesJSON, &event.Capacity, &event.CreatorID, &event.ImageURL, &hostName, &event.Participants); err != nil {
			return nil, fmt.Errorf("failed to scan searched event: %w", err)
		}

		if err := json.Unmarshal([]byte(rulesJSON), &event.Rules); err != nil {
			return nil, fmt.Errorf("failed to unmarshal rules for searched event ID %d: %w", event.ID, err)
		}
		event.HostName = hostName
		events = append(events, event)
	}

	return events, nil
}
