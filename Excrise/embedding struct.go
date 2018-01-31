package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First   string
	License bool
}

func main() {

	p1 := DoubleZero{
		Person: Person{
			"Abhishek",
			"Singh",
			25,
		},
		First:   "Double Zero Abhishek",
		License: true,
	}

	p2 := DoubleZero{
		Person: Person{
			"Bobby",
			"Singh",
			25,
		},
		First:   "Double Zero Bobby",
		License: true,
	}

	fmt.Println(p1.Person.First, p1.Person.Last, p1.Person.Age, p1.First, p1.License)
	fmt.Println(p2.Person.First, p2.Person.Last, p2.Person.Age, p2.First, p2.License)
}
