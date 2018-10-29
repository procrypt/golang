package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)
	goroutines := 10

	for i := 0; i < goroutines; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
}
