package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB // global variable

func InitDB() {
	var err error
        DB, err = sql.Open("sqlite3", "api.db") // open connection
     
        if err != nil {
            panic("Could not connect to database.")
        }
     
        DB.SetMaxOpenConns(10) // pool of ongoing connections
        DB.SetMaxIdleConns(5) // how many connections we want to keep open if no one is using them
     
        createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  title TEXT NOT NULL,
	  description TEXT NOT NULL,
	  location TEXT NOT NULL,
	  dateTime DATETIME NOT NULL,
	  user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table!")
	}
}