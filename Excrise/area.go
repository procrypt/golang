package main

import (
	"math"
	"fmt"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	lenght float64
	width float64
}

func(c Circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func(c Circle) perimeter() float64 {
	return 2*math.Pi*c.radius
}

func (r Rectangle) area() float64 {
	return r.width*r.lenght
}

func (r Rectangle) perimeter() float64 {
	return 2*(r.lenght+r.lenght)
}

func main() {
	c := Circle{10}
	r := Rectangle{5,10}
	fmt.Println(c.area(), c.perimeter())
	fmt.Println(r.area(), r.perimeter())
}
