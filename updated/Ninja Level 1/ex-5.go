package main

import "fmt"

type abhi int
var x abhi
var z int

func main()  {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	z = int(x)
	fmt.Println(z)
	fmt.Printf("%T\n",z)

}