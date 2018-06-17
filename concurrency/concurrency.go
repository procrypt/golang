package main

import (
	"fmt"
	"sync"
)


var wg sync.WaitGroup

func foo() {
	for i:=1; i<=1000; i++ {
		fmt.Printf("Foo: %d\n", i)
	}
	wg.Done() // Notify the main func that it has done its execution
}

func bar() {
	for i:=1; i<=1000; i++ {
		fmt.Printf("Bar: %d\n", i)
	}
	wg.Done()
}

func main() {
	wg.Add(2) // Add two wait groups foo and bar, once the function is done with its execution, delta decreases by 1
	go foo()
	go bar()
	wg.Wait()
}
