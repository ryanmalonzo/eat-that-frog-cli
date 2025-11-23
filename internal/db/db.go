// Package db provides database connection and utility functions.
package db

import (
	"database/sql"
	"time"

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

func AddCandidate(task string) error {
	_, err := dbPool.Exec("INSERT INTO candidates (task) VALUES (?)", task)
	return err
}

func GetAllCandidates() ([]string, error) {
	rows, err := dbPool.Query("SELECT task FROM candidates ORDER BY created_at")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var candidates []string
	for rows.Next() {
		var task string
		if err := rows.Scan(&task); err != nil {
			return nil, err
		}
		candidates = append(candidates, task)
	}

	return candidates, nil
}

func DeleteAllCandidates() error {
	_, err := dbPool.Exec("DELETE FROM candidates")
	return err
}

func CountCandidates() (int, error) {
	var count int
	err := dbPool.QueryRow("SELECT COUNT(*) FROM candidates").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func PickCandidate(index int) error {
	tx, err := dbPool.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var id int
	var task string
	err = tx.QueryRow("SELECT id, task FROM candidates ORDER BY created_at LIMIT 1 OFFSET ?", index).Scan(&id, &task)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO frogs (date, task, status) VALUES (CURRENT_DATE, ?, 'pending')", task)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM candidates WHERE id = ?", id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func GetTodayFrog() (string, error) {
	var task string
	err := dbPool.QueryRow("SELECT task FROM frogs WHERE date = CURRENT_DATE").Scan(&task)
	if err == sql.ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return task, nil
}

func updateFrogStatus(task string, status string, timestamp time.Time) error {
	_, err := dbPool.Exec("UPDATE frogs SET status = ?, updated_at = ? WHERE date = CURRENT_DATE AND task = ?", status, timestamp, task)
	return err
}

func MarkFrogAsDone(task string) error {
	return updateFrogStatus(task, "done", time.Now())
}

func SkipTodayFrog(task string) error {
	return updateFrogStatus(task, "skip", time.Now())
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
