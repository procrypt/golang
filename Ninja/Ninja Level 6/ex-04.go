package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Printf("Hi my name is %s and my age is %d.\n", p.first, p.age)
}

func main() {
	p1 := person{
		first: "Abhishek",
		last:  "Singh",
		age: 25,
	}
	p1.speak()
}