package main

import "fmt"

func foo(n ...int) {
	fmt.Println("This is correct!!!")
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aslice := []int{1, 2, 3, 4}
	foo(aslice...)
	foo()
}
