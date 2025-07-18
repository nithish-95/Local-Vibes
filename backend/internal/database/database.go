package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/nithish-95/Local-Vibes/backend/internal/models"
)

var DB *gorm.DB

func ConnectDB(dbPath string) (*gorm.DB, error) {
	absPath, err := filepath.Abs(dbPath)
	if err != nil {
		return nil, fmt.Errorf("Error getting absolute path for %s: %v", dbPath, err)
	}

	dir := filepath.Dir(absPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Printf("Database directory %s does not exist. Creating...", dir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("Error creating database directory %s: %v", dir, err)
		}
		log.Printf("Database directory %s created successfully.", dir)
	} else if err != nil {
		return nil, fmt.Errorf("Error checking database directory %s: %v", dir, err)
	} else {
		log.Printf("Database directory %s already exists.", dir)
	}

	if _, err := os.Stat(absPath); err == nil {
		log.Printf("Database file %s already exists.", absPath)
	} else if os.IsNotExist(err) {
		log.Printf("Database file %s does not exist. It will be created on open.", absPath)
	} else {
		return nil, fmt.Errorf("Error checking database file %s: %v", absPath, err)
	}

	log.Printf("Attempting to open database at: %s", absPath)
	db, err := gorm.Open(sqlite.Open(absPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	// AutoMigrate will create/update tables based on GORM models
	err = db.AutoMigrate(&models.User{}, &models.Event{})
	if err != nil {
		return nil, fmt.Errorf("Failed to auto migrate database: %w", err)
	}

	// Create event_participants table manually as GORM doesn't handle many-to-many with custom join table easily
	createEventParticipantsTableSQL := `CREATE TABLE IF NOT EXISTS event_participants (
		"event_id" INTEGER NOT NULL,
		"user_id" INTEGER NOT NULL,
		PRIMARY KEY (event_id, user_id),
		FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	err = db.Exec(createEventParticipantsTableSQL).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to create event_participants table: %w", err)
	}

	// Create FTS5 virtual table for events
	createEventsFTS := `CREATE VIRTUAL TABLE IF NOT EXISTS events_fts USING fts5(title, description, location, rules, content='events', content_rowid='id');`
	err = db.Exec(createEventsFTS).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to create FTS5 table: %w", err)
	}

	// Triggers to keep events_fts in sync with events table
	createEventsFTSInsertTrigger := `
	CREATE TRIGGER IF NOT EXISTS events_ai AFTER INSERT ON events BEGIN
		INSERT INTO events_fts(rowid, title, description, location, rules) VALUES (new.id, new.title, new.description, new.location, new.rules);
	END;`
	err = db.Exec(createEventsFTSInsertTrigger).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to create FTS5 insert trigger: %w", err)
	}

	createEventsFTSUpdateTrigger := `
	CREATE TRIGGER IF NOT EXISTS events_au AFTER UPDATE ON events BEGIN
		INSERT INTO events_fts(events_fts, rowid, title, description, location, rules) VALUES('delete', old.id, old.title, old.description, old.location, old.rules);
		INSERT INTO events_fts(rowid, title, description, location, rules) VALUES (new.id, new.title, new.description, new.location, new.rules);
	END;`
	err = db.Exec(createEventsFTSUpdateTrigger).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to create FTS5 update trigger: %w", err)
	}

	createEventsFTSDeleteTrigger := `
	CREATE TRIGGER IF NOT EXISTS events_ad AFTER DELETE ON events BEGIN
		INSERT INTO events_fts(events_fts, rowid, title, description, location, rules) VALUES('delete', old.id, old.title, old.description, old.location, old.rules);
	END;`
	err = db.Exec(createEventsFTSDeleteTrigger).Error
	if err != nil {
		return nil, fmt.Errorf("Failed to create FTS5 delete trigger: %w", err)
	}

	return db, nil
}

func InitDB() {
	var err error
	dbPath := "./database/app.db"
	DB, err = ConnectDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
}
