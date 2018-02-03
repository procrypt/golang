package main

import "fmt"

type Person struct {
	First string
	Last string
	Age int
}


type DoubleZero struct {
	Person
	First string
	License bool
}

func (p Person) Greetings() string {
	return "Hello Person"
}

func (d DoubleZero) Greetings() string {
	return "Hello DoubleZero"
}


func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "Abhishek",
			Last: "Singh",
			Age: 25,
		},
		First: "abhishek",
		License: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Bobby",
			Last: "Singh",
			Age: 25,
		},
		First: "bobby",
		License: true,
	}

	fmt.Println(p1.Greetings())
	fmt.Println(p2.Person.Greetings())
}