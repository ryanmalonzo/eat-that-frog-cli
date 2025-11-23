// Package db provides database connection and utility functions.
package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

const DBFileName = "frog.db"

var dbPool *sql.DB

func Init() error {
	var err error
	dbPool, err = sql.Open("sqlite", DBFileName)
	if err != nil {
		return err
	}

	return createTables(dbPool)
}

func Close() error {
	if dbPool != nil {
		return dbPool.Close()
	}
	return nil
}

func createTables(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS frogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TIMESTAMP NOT NULL UNIQUE,
		task TEXT NOT NULL,
		status TEXT NOT NULL CHECK(status IN ('pending', 'done', 'skip')),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS candidates (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_frogs_date ON frogs(date);
	CREATE INDEX IF NOT EXISTS idx_frogs_status ON frogs(status);
	`

	_, err := db.Exec(schema)
	return err
}
