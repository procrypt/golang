package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
