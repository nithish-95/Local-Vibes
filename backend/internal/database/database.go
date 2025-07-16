package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	dbPath := "./database/app.db"
	absPath, err := filepath.Abs(dbPath)
	if err != nil {
		log.Fatalf("Error getting absolute path for %s: %v", dbPath, err)
	}

	dir := filepath.Dir(absPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Printf("Database directory %s does not exist. Creating...", dir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Error creating database directory %s: %v", dir, err)
		}
		log.Printf("Database directory %s created successfully.", dir)
	} else if err != nil {
		log.Fatalf("Error checking database directory %s: %v", dir, err)
	} else {
		log.Printf("Database directory %s already exists.", dir)
	}

	if _, err := os.Stat(absPath); err == nil {
		log.Printf("Database file %s already exists.", absPath)
	} else if os.IsNotExist(err) {
		log.Printf("Database file %s does not exist. It will be created on open.", absPath)
	} else {
		log.Fatalf("Error checking database file %s: %v", absPath, err)
	}

	log.Printf("Attempting to open database at: %s", absPath)
	DB, err = sql.Open("sqlite3", absPath)
	if err != nil {
		log.Fatal(err)
	}

	createUsersTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT NOT NULL UNIQUE,
		"password" TEXT NOT NULL
	);`

	_, err = DB.Exec(createUsersTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createEventsTableSQL := `CREATE TABLE IF NOT EXISTS events (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT NOT NULL,
		"description" TEXT,
		"date" TEXT NOT NULL,
		"time" TEXT NOT NULL,
		"location" TEXT NOT NULL,
		"rules" TEXT,
		"capacity" INTEGER NOT NULL DEFAULT 0,
		"image_url" TEXT,
		"creator_id" INTEGER,
		FOREIGN KEY (creator_id) REFERENCES users(id)
	);`

	_, err = DB.Exec(createEventsTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createEventParticipantsTableSQL := `CREATE TABLE IF NOT EXISTS event_participants (
		"event_id" INTEGER NOT NULL,
		"user_id" INTEGER NOT NULL,
		PRIMARY KEY (event_id, user_id),
		FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	_, err = DB.Exec(createEventParticipantsTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Create FTS5 virtual table for events
	createEventsFTS := `CREATE VIRTUAL TABLE IF NOT EXISTS events_fts USING fts5(title, description, location, rules, content='events', content_rowid='id');`
	_, err = DB.Exec(createEventsFTS)
	if err != nil {
		log.Fatal(err)
	}

	// Triggers to keep events_fts in sync with events table
	createEventsFTSInsertTrigger := `
	CREATE TRIGGER IF NOT EXISTS events_ai AFTER INSERT ON events BEGIN
		INSERT INTO events_fts(rowid, title, description, location, rules) VALUES (new.id, new.title, new.description, new.location, new.rules);
	END;`
	_, err = DB.Exec(createEventsFTSInsertTrigger)
	if err != nil {
		log.Fatal(err)
	}

	createEventsFTSUpdateTrigger := `
	CREATE TRIGGER IF NOT EXISTS events_au AFTER UPDATE ON events BEGIN
		INSERT INTO events_fts(events_fts, rowid, title, description, location, rules) VALUES('delete', old.id, old.title, old.description, old.location, old.rules);
		INSERT INTO events_fts(rowid, title, description, location, rules) VALUES (new.id, new.title, new.description, new.location, new.rules);
	END;`
	_, err = DB.Exec(createEventsFTSUpdateTrigger)
	if err != nil {
		log.Fatal(err)
	}

	createEventsFTSDeleteTrigger := `
	CREATE TRIGGER IF NOT EXISTS events_ad AFTER DELETE ON events BEGIN
		INSERT INTO events_fts(events_fts, rowid, title, description, location, rules) VALUES('delete', old.id, old.title, old.description, old.location, old.rules);
	END;`
	_, err = DB.Exec(createEventsFTSDeleteTrigger)
	if err != nil {
		log.Fatal(err)
	}
}
