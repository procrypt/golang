package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("bar", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("foo", i)
		}
		wg.Done()
	}()

	wg.Wait()
}
