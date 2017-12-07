package main

import (
	"fmt"
)

func greatest(number ...int) int {
	var largest int
	for _,v := range number {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	fmt.Println(greatest(4, 5, 6, 1, 2, 8, 7, 2, 3, 50))
}
