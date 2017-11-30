package main

import "fmt"

func main() {

	var a int
	var b int
	var c *int

	a = 1
	b = a
	c = &a

	*c = 21

	fmt.Printf("Values:\n a = %v\n b = %v\n c = %v\n", a, b, c)
	fmt.Printf("Values:\n a = %v\n b = %v\n c = %v\n", a, b, c)
	fmt.Printf("Values:\n a = %v\n b = %v\n c = %v\n", a, b, *c)
}
