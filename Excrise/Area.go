package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width float64
}


func (c Circle) Area() float64 {
	return 2*3.14*c.radius
}

func (c Circle) Perimeter() float64 {
	return 3.14*math.Sqrt(c.radius)
}

func (r Rectangle) Area() float64 {
	return r.length*r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2*(r.length+r.width)
}

func main() {
	c := Circle{
		radius: 10,
	}
	r := Rectangle{
		length: 5,
		width: 8,

	}

	fmt.Println("Area of Circle is: ", c.Area())
	fmt.Println("Area of Rectangle is: ", r.Area())

	fmt.Println("Perimeter of Circle is: ", c.Perimeter())
	fmt.Println("Perimeter of Rectangle is: ", r.Perimeter())
}