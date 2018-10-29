package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c <-chan int, q chan int) {
	for i := 0; i < 100; i++ {
		select {
		case msg1 := <-c:
			q <- msg1
		case msg2 := <-q:
			fmt.Println(msg2)
		}
	}
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
