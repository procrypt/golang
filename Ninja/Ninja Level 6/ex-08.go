package main

import "fmt"

func foo() func() {
	total := 0
	return func() {
		for i:=0; i<20; i++ {
			total += i
		}
		fmt.Println(total)
	}
}

func main() {
	a := foo()
	a()
}