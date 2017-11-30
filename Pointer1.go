package main

import "fmt"

func zero (z *int) {
	*z = 10
}

func main() {
	x := 20
	fmt.Println(&x)
	zero(&x)
	fmt.Println(&x)
	fmt.Println(x)
}
