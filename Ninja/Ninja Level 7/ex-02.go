package main

import "fmt"

type person struct {
	first string
}

func changeMe(p *person) {
	p.first = "Bobby"

}

func main() {
	p1 := person{
		first: "Abhishek",
	}
	fmt.Println(p1.first)
	changeMe(&p1)
	fmt.Println(p1.first)
}
