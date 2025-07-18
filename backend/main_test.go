package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/nithish-95/Local-Vibes/backend/internal/database"
	"github.com/nithish-95/Local-Vibes/backend/internal/models"
	"github.com/nithish-95/Local-Vibes/backend/internal/routes"
)

// Helper function to create a temporary database for testing
func setupTestDB(t *testing.T) string {
	log.SetOutput(io.Discard) // Suppress application logs during test
	tempDir, err := os.MkdirTemp("", "testdb")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	dbPath := filepath.Join(tempDir, "test.db")
	// log.Printf("Using temporary database at: %s", dbPath) // Keep this commented or remove if you want absolutely no test setup logs

	// Set SESSION_SECRET_KEY for the test environment
	os.Setenv("SESSION_SECRET_KEY", "test-secret-key-for-session-management-123456")

	// Initialize the database with the test path
	db, err := database.ConnectDB(dbPath) // Assuming ConnectDB is now exported
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	database.DB = db // Assign to global DB for handlers/services

	return tempDir
}

// Helper function to clean up the temporary database
func teardownTestDB(tempDir string) {
	if database.DB != nil {
		sqlDB, err := database.DB.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
	os.RemoveAll(tempDir)
	os.Unsetenv("SESSION_SECRET_KEY")
	log.SetOutput(os.Stderr) // Restore default logging
}

func TestFullFlow(t *testing.T) {
	tempDir := setupTestDB(t)
	defer teardownTestDB(tempDir)

	router := routes.SetupRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	client := ts.Client()

	var hostedEvents []models.Event
	var found bool

	// --- User 1 (Creator) Flow ---

	// 1. Register User 1
	registerPayload1 := map[string]string{
		"username": "testuser_flow1",
		"password": "testpassword_flow1",
	}
	registerBody1, _ := json.Marshal(registerPayload1)
	resp, err := client.Post(fmt.Sprintf("%s/api/register", ts.URL), "application/json", bytes.NewBuffer(registerBody1))
	if err != nil {
		t.Fatalf("Failed to register user 1: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Register user 1 failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	resp.Body.Close()
	t.Log("User 1 registered successfully.")

	// 2. Login User 1 and get session cookie
	loginPayload1 := map[string]string{
		"username": "testuser_flow1",
		"password": "testpassword_flow1",
	}
	loginBody1, _ := json.Marshal(loginPayload1)
	resp, err = client.Post(fmt.Sprintf("%s/api/login", ts.URL), "application/json", bytes.NewBuffer(loginBody1))
	if err != nil {
		t.Fatalf("Failed to login user 1: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Login user 1 failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	var sessionCookie1 *http.Cookie
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "session-name" {
			sessionCookie1 = cookie
			break
		}
	}
	if sessionCookie1 == nil {
		t.Fatal("Session cookie for user 1 not found after login")
	}
	resp.Body.Close()
	t.Log("User 1 logged in and session cookie obtained.")

	// 2.1 Get Current User
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/user", ts.URL), nil)
	req.AddCookie(sessionCookie1)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to get current user: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Get current user failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	var currentUser models.User
	json.NewDecoder(resp.Body).Decode(&currentUser)
	resp.Body.Close()
	if currentUser.Username != "testuser_flow1" {
		t.Fatalf("Expected username testuser_flow1, got %s", currentUser.Username)
	}
	t.Log("Current user details fetched successfully.")

	// 3. Create Event by User 1
	createEventPayload := map[string]interface{}{
		"title":       "Test Event from Flow",
		"description": "This is a test event created via the full flow test.",
		"date":        "2025-12-25",
		"time":        "10:00",
		"location":    "Test Location",
		"rules":       "[\"Rule 1\", \"Rule 2\"]", // Send as JSON string
		"capacity":    10,
		"image_url":   "",
	}
	createEventBody, _ := json.Marshal(createEventPayload)
	req, _ = http.NewRequest("POST", fmt.Sprintf("%s/api/events", ts.URL), bytes.NewBuffer(createEventBody))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(sessionCookie1)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to create event: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Create event failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	var createEventResponse map[string]uint
	json.NewDecoder(resp.Body).Decode(&createEventResponse)
	createdEventID := createEventResponse["id"]
	resp.Body.Close()
	t.Logf("Event created successfully with ID: %d", createdEventID)

	// 3.1 Get All Events
	req, _ = http.NewRequest("GET", fmt.Sprintf("%s/api/events", ts.URL), nil)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to get all events: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Get all events failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	var allEvents []models.Event
	json.NewDecoder(resp.Body).Decode(&allEvents)
	resp.Body.Close()
	if len(allEvents) == 0 {
		t.Fatal("Expected to get all events, but none found.")
	}
	// Verify the newly created event is in the list of all events
	found = false
	for _, event := range allEvents {
		if event.ID == createdEventID {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Created event (ID %d) not found in all events list.", createdEventID)
	}
	t.Log("All events fetched successfully.")

	// 3.2 Get Event by ID
	req, _ = http.NewRequest("GET", fmt.Sprintf("%s/api/events/%d", ts.URL, createdEventID), nil)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to get event by ID: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Get event by ID failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	var fetchedEvent models.Event
	json.NewDecoder(resp.Body).Decode(&fetchedEvent)
	resp.Body.Close()
	if fetchedEvent.ID != createdEventID || fetchedEvent.Title != "Test Event from Flow" {
		t.Fatalf("Fetched event by ID mismatch. Expected ID %d, got %d. Expected title \"Test Event from Flow\", got \"%s\"", createdEventID, fetchedEvent.ID, fetchedEvent.Title)
	}
	t.Logf("Event ID %d fetched by ID successfully.", createdEventID)

	// 3.3 Update Event
	updateEventPayload := map[string]interface{}{
		"id":          createdEventID,
		"title":       "Updated Test Event",
		"description": "This event has been updated by the flow test.",
		"date":        "2026-01-01",
		"time":        "12:00",
		"location":    "Updated Location",
		"rules":       "[\"Updated Rule 1\", \"Updated Rule 2\"]",
		"capacity":    20,
		"image_url":   "http://example.com/updated.jpg",
	}
	updateEventBody, _ := json.Marshal(updateEventPayload)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("%s/api/events/%d", ts.URL, createdEventID), bytes.NewBuffer(updateEventBody))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(sessionCookie1)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to update event: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Update event failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	resp.Body.Close()
	t.Logf("Event ID %d updated successfully.", createdEventID)

	// 4. Get Hosted Events by User 1 and verify (after update)
	req, _ = http.NewRequest("GET", fmt.Sprintf("%s/api/events/hosted", ts.URL), nil)
	req.AddCookie(sessionCookie1)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to get hosted events after update: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Get hosted events after update failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	json.NewDecoder(resp.Body).Decode(&hostedEvents)
	resp.Body.Close()

	if len(hostedEvents) == 0 {
		t.Fatal("No hosted events found for user 1 after update, expected at least one.")
	}

	found = false
	for _, event := range hostedEvents {
		if event.ID == createdEventID && event.Title == "Updated Test Event" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Updated event (ID %d) not found in user 1's hosted events list.", createdEventID)
	}
	t.Log("Updated event found in user 1's hosted events list.")

	// --- User 2 (Participant) Flow ---

	// 5. Register User 2
	registerPayload2 := map[string]string{
		"username": "testuser_flow2",
		"password": "testpassword_flow2",
	}
	registerBody2, _ := json.Marshal(registerPayload2)
	resp, err = client.Post(fmt.Sprintf("%s/api/register", ts.URL), "application/json", bytes.NewBuffer(registerBody2))
	if err != nil {
		t.Fatalf("Failed to register user 2: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Register user 2 failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	resp.Body.Close()
	t.Log("User 2 registered successfully.")

	// 6. Login User 2 and get session cookie
	loginPayload2 := map[string]string{
		"username": "testuser_flow2",
		"password": "testpassword_flow2",
	}
	loginBody2, _ := json.Marshal(loginPayload2)
	resp, err = client.Post(fmt.Sprintf("%s/api/login", ts.URL), "application/json", bytes.NewBuffer(loginBody2))
	if err != nil {
		t.Fatalf("Failed to login user 2: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Login user 2 failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	var sessionCookie2 *http.Cookie
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "session-name" {
			sessionCookie2 = cookie
			break
		}
	}
	if sessionCookie2 == nil {
		t.Fatal("Session cookie for user 2 not found after login")
	}
	resp.Body.Close()
	t.Log("User 2 logged in and session cookie obtained.")

	// 6.1 Get Available Events for User 2
	req, _ = http.NewRequest("GET", fmt.Sprintf("%s/api/events/available", ts.URL), nil)
	req.AddCookie(sessionCookie2)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to get available events for user 2: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Get available events for user 2 failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	var availableEvents []models.Event
	json.NewDecoder(resp.Body).Decode(&availableEvents)
	resp.Body.Close()
	// Verify the event created by user 1 is available to user 2
	found = false
	for _, event := range availableEvents {
		if event.ID == createdEventID {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Event ID %d not found in user 2's available events list.", createdEventID)
	}
	t.Logf("Event ID %d found in user 2's available events list.", createdEventID)

	// 6.2 Search Events for User 2
	req, _ = http.NewRequest("GET", fmt.Sprintf("%s/api/events/search?q=Updated&startDate=2026-01-01", ts.URL), nil)
	req.AddCookie(sessionCookie2)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to search events for user 2: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Search events for user 2 failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	var searchResults []models.Event
	json.NewDecoder(resp.Body).Decode(&searchResults)
	resp.Body.Close()
	// Verify the updated event is found in search results
	found = false
	for _, event := range searchResults {
		if event.ID == createdEventID && event.Title == "Updated Test Event" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Updated event (ID %d) not found in user 2's search results.", createdEventID)
	}
	t.Logf("Updated event (ID %d) found in user 2's search results.", createdEventID)

	// 7. User 2 joins the event created by User 1
	req, _ = http.NewRequest("POST", fmt.Sprintf("%s/api/events/%d/join", ts.URL, createdEventID), nil)
	req.AddCookie(sessionCookie2)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed for user 2 to join event: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("User 2 join event failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	resp.Body.Close()
	t.Logf("User 2 successfully joined event ID: %d", createdEventID)

	// 8. User 2 checks if joined to the event
	req, _ = http.NewRequest("GET", fmt.Sprintf("%s/api/events/%d/is-joined", ts.URL, createdEventID), nil)
	req.AddCookie(sessionCookie2)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed for user 2 to check if joined: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("User 2 check if joined failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	var isJoinedResponse map[string]bool
	json.NewDecoder(resp.Body).Decode(&isJoinedResponse)
	resp.Body.Close()
	if !isJoinedResponse["is_joined"] {
		t.Fatalf("User 2 expected to be joined to event %d, but is_joined is false.", createdEventID)
	}
	t.Logf("User 2 confirmed as joined to event ID: %d", createdEventID)

	// 9. User 2 gets joined events and verifies the event is listed
	req, _ = http.NewRequest("GET", fmt.Sprintf("%s/api/events/joined", ts.URL), nil)
	req.AddCookie(sessionCookie2)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed for user 2 to get joined events: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("User 2 get joined events failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	var joinedEvents []models.Event
	json.NewDecoder(resp.Body).Decode(&joinedEvents)
	resp.Body.Close()

	if len(joinedEvents) == 0 {
		t.Fatalf("User 2 expected to have joined events, but none found.")
	}
	found = false
	for _, event := range joinedEvents {
		if event.ID == createdEventID {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Event ID %d not found in user 2's joined events list.", createdEventID)
	}
	t.Logf("Event ID %d found in user 2's joined events list.", createdEventID)

	// 10. User 2 leaves the event
	req, _ = http.NewRequest("POST", fmt.Sprintf("%s/api/events/%d/leave", ts.URL, createdEventID), nil)
	req.AddCookie(sessionCookie2)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed for user 2 to leave event: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("User 2 leave event failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	resp.Body.Close()
	t.Logf("User 2 successfully left event ID: %d", createdEventID)

	// 11. User 2 checks if joined to the event (should be false)
	req, _ = http.NewRequest("GET", fmt.Sprintf("%s/api/events/%d/is-joined", ts.URL, createdEventID), nil)
	req.AddCookie(sessionCookie2)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed for user 2 to check if joined after leaving: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("User 2 check if joined after leaving failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	json.NewDecoder(resp.Body).Decode(&isJoinedResponse)
	resp.Body.Close()
	if isJoinedResponse["is_joined"] {
		t.Fatalf("User 2 expected not to be joined to event %d after leaving, but is_joined is true.", createdEventID)
	}
	t.Logf("User 2 confirmed as not joined to event ID: %d after leaving.", createdEventID)

	// 12. User 2 gets joined events and verifies the event is NOT listed
	req, _ = http.NewRequest("GET", fmt.Sprintf("%s/api/events/joined", ts.URL), nil)
	req.AddCookie(sessionCookie2)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed for user 2 to get joined events after leaving: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("User 2 get joined events after leaving failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	json.NewDecoder(resp.Body).Decode(&joinedEvents)
	resp.Body.Close()

	found = false
	for _, event := range joinedEvents {
		if event.ID == createdEventID {
			found = true
			break
		}
	}
	if found {
		t.Fatalf("Event ID %d found in user 2's joined events list after leaving, but should not be.", createdEventID)
	}
	t.Logf("Event ID %d not found in user 2's joined events list after leaving.", createdEventID)

	// --- Cleanup Flow (User 1) ---

	// 13. User 1 Logout
	req, _ = http.NewRequest("POST", fmt.Sprintf("%s/api/logout", ts.URL), nil)
	req.AddCookie(sessionCookie1)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed for user 1 to logout: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("User 1 logout failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	resp.Body.Close()
	t.Log("User 1 logged out successfully.")

	// 14. Delete Event by User 1
	req, _ = http.NewRequest("DELETE", fmt.Sprintf("%s/api/events/%d", ts.URL, createdEventID), nil)
	req.AddCookie(sessionCookie1)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to delete event: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Delete event failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}
	resp.Body.Close()
	t.Logf("Event ID %d deleted successfully.", createdEventID)
}
