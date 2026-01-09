package repository

import "database/sql"

func SaveTransaction(db *sql.DB, stan, rrn string, amount int64) error {
	_, err := db.Exec(
		"INSERT INTO transactions (stan, rrn, amount, status) VALUES (?, ?, ?, ?)",
		stan, rrn, amount, "PENDING",
	)
	return err
}
