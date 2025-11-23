// Package db provides database connection and utility functions.
package db

import (
	"database/sql"
	"time"
)

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

