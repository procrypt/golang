package main

import "fmt"

type Person struct {
	first string
	last string
	age int
}

func(p Person) fullname() string {
	return  p.first + p.last
}


func main() {

	p1 := Person{"Abhishek", " Singh", 25}
	p2 := Person{"Bobby", " Singh", 25}

	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())
}