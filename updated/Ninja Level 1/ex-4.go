package main

import "fmt"

type abhi int
var x abhi

func main()  {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}