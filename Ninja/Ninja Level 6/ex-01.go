package main

import "fmt"

func foo() int {
	x := 40
	return x
}

func bar() (int, string) {
	x := 43
	y := "Abhishek"
	return x,y
}

func main()  {
	n := foo()
	x,s := bar()
	fmt.Println(n,x,s)
}