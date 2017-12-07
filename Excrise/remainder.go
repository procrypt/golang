package main

import "fmt"

func main() {
	var large, small int
	fmt.Println("Enter large number:")
	fmt.Scan(&large)
	fmt.Println("Enter small number:")
	fmt.Scan(&small)
	result := large%small
	fmt.Print("Remainder: ", result)
}
