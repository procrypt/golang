package main

import "fmt"

type Person struct {
	First string
	Last string
	age int
}

func main() {
	p1 := Person{"Abhishek", "Singh", 25}
	fmt.Println(p1.First, p1.Last, p1.age)
}
