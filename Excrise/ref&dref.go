package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a  // b is referencing a memory address

	fmt.Println(b)
	fmt.Println(*b) // dereferencing the memory address, which gives the value at that memory address

	*b = 42 // Change the value at address *b to 42

	fmt.Println(a)

}
