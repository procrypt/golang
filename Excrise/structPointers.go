package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	p1 := &Person{"Abhishek", 25}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
}
