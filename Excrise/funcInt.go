package main

import "fmt"

func todd(n int) (int, bool) {
	result := n/2
	if n%2 == 0 {
		return result, true
	} else {
		return result, false
	}
}

func main() {
	fmt.Println(todd(2))
}
