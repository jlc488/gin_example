package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite", "api.db")
	if err != nil {
		panic("Unable to connect")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	DB = db

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   name TEXT NOT NULL,
	   description TEXT NOT NULL,
	   dateTime DATETIME NOT NULL,
	   user_id INTEGER,
	   FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create events table")
	}

	createREgistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    event_id INTEGER,
	    user_id INTEGER,
	    FOREIGN KEY(event_id) REFERENCES events(id),
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createREgistrationTable)
	if err != nil {
		panic("could not create registrations table")
	}

}
