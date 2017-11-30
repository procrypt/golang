package main

import "fmt"

func main() {
	for i := 0; i <= 5000; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
}
