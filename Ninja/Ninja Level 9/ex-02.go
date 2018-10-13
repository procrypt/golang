package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() bool {
	fmt.Println("The person speaks!!")
	return true
}

type human interface {
	speak() bool
}

func saysomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "Abhishek",
		last:  "Singh",
		age:   25,
	}

	saysomething(&p)
}
