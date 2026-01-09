package main

import (
	"fmt"
	"net"

	"mini-payment-switch/bank-simulator/iso"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		panic(err)
	}

	msg := iso.Build0200()
	conn.Write([]byte(msg))

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	fmt.Println("Bank received:", string(buf[:n]))
}
