package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	fmt.Println("Payment Switch running on :5000")

	for {
		conn, _ := ln.Accept()
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	msg := string(buf[:n])
	resp := processISO(msg)

	conn.Write([]byte(resp))
}
