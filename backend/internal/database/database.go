package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "../database/app.db")
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
}