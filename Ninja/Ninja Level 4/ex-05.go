package main

import "fmt"

func main()  {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// [42, 43, 44, 48, 49, 50, 51]

	y := append(x[0:3], x[6:]...)
	fmt.Println(y)
}
