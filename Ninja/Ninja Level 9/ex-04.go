package main

import (
	"fmt"
	"sync"
)

func main() {

	var counter int

	var mu sync.Mutex

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
