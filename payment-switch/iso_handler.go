package main

import "fmt"

func processISO(msg string) string {
	fmt.Println("Processing ISO:", msg)

	// TODO later:
	// validate MAC
	// persist transaction
	// call wallet-service

	return "0210|RC=00"
}
