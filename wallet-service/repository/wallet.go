package repository

import "database/sql"

func DebitWallet(tx *sql.Tx, walletID string, amount int) error {
	_, err := tx.Exec(
		"UPDATE wallets SET balance = balance - ? WHERE id = ?",
		amount, walletID,
	)
	return err
}

func CreditWallet(tx *sql.Tx, walletID string, amount int) error {
	_, err := tx.Exec(
		"UPDATE wallets SET balance = balance + ? WHERE id = ?",
		amount, walletID,
	)
	return err
}
