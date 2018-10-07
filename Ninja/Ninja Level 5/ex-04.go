package main

import "fmt"

func main() {
	s := struct {
		first string
		last string
		age int
	}{
		first: "Abhishek",
		last: "Singh",
		age: 25,
	}
	fmt.Println(s)
}
