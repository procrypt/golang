package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var counter int

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
