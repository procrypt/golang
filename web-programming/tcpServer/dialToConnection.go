package main

import (
	"net"
	"log"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err!=nil{
		log.Fatalln(err)
	}
	defer conn.Close()
	data := []byte("I dialed you!")
	_, err = conn.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
}
