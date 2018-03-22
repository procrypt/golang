package main

import (
	"net"
	"io"
	"time"
	"log"
	"fmt"
	"os"
	"strings"
	)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func handleConn(c net.Conn) {
	args := os.Args[1:]
	tz := strings.Split(args[1], ":")
	if tz[0] == "Tokyo" {
		err := os.Setenv("TZ", "Asia/Tokyo")
		check(err)
	} else {
		if tz[0] == "NewYork" {
			err := os.Setenv("TZ", "US/Eastern")
			check(err)
		} else {
			if tz[0] == "London" {
				err := os.Setenv("TZ", "Europe/London")
				check(err)
			}
		}
		defer c.Close()
		for {
			_, err := io.WriteString(c, time.Now().Format("Jan Mon 15:04:05\n"))
			if err != nil {
				log.Print(err)
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
		a := os.Args[1]
		p := strings.Split(a, ":")
		port := p[1]
		listener, err := net.Listen("tcp", "localhost:"+port)
		if err != nil {
			log.Fatal(err)
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
			}
			fmt.Println(conn.RemoteAddr())
			go handleConn(conn)
		}
	}