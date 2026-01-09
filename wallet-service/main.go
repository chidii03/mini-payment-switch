package main

import (
	"fmt"
	"log"

	"mini-payment-switch/wallet-service/db"
	"mini-payment-switch/wallet-service/repository"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	tx, err := database.Begin()
	if err != nil {
		log.Fatal(err)
	}

	ref := "STAN123456-RRN654321"

	err = repository.InsertLedger(tx, "wallet-1", 1000, "DEBIT", ref)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = repository.DebitWallet(tx, "wallet-1", 1000)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
	fmt.Println("Wallet debited successfully")
}
