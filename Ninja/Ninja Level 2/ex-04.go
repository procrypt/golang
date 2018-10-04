package main

import "fmt"

func main() {

	x := 10
	fmt.Printf("%d\t%b\t%#x\t", x,x,x)
	fmt.Println("")
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\t", y,y,y)
}
