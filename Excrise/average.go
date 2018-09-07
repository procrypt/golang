package main

import "fmt"

func average(n ...float64) float64 {
	var total float64
	for _, v := range n {
		total += v
	}
	return total/float64(len(n))
}

func main() {
	result := average(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(result)
}
