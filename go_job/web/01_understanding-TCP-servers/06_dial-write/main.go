package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err !=nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
