package main

import "fmt"

func foo(x ...int) int {
	total := 0
	for _,v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _,v := range x {
		total += v
	}
	return total
}

func main() {
	x := []int{1,2,3,4,5,6,7,8,9}
	s1 := foo(x...)
	s2 := bar(x)
	fmt.Println(s1)
	fmt.Println(s2)
}