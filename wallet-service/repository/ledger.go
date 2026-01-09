package repository

import "database/sql"

func InsertLedger(tx *sql.Tx, walletID string, amount int, entryType, ref string) error {
	_, err := tx.Exec(
		`INSERT INTO ledger(wallet_id, amount, type, reference)
		 VALUES(?,?,?,?)`,
		walletID, amount, entryType, ref,
	)
	return err
}
