package main

import "fmt"

func main() {
	a:= make([]int, 50, 100)

	for i := 0; i < 50; i++ {
		a[i] = i
		fmt.Println(len(a))
	}

	for j := 50; j < 500; j++ {
		a = append(a, j)
		fmt.Println(cap(a))
	}
}
