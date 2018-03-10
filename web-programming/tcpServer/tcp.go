package main

import (
	"net"
	"log"
	"io"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err!=nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err!= nil {
			log.Fatalln(err)
		}

		io.WriteString(conn, "Hello from TCP Server in Go!\n")
		fmt.Fprintln(conn, "How are you doing today?")
		fmt.Fprintf(conn, "%v", "Hope you are doing good!\n")

		conn.Close()
	}
}
