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

type Shape interface {
	area() float64
	perimeter() float64
}

func (c Circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func (c Circle) perimeter() float64 {
	return 2*math.Pi*c.radius
}

func (r Rectangle) area() float64 {
	return r.lenght*r.width
}

func (r Rectangle) perimeter() float64 {
	return 2*(r.width+r.lenght)
}

func info(s Shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}


func main() {
	c := Circle{10}
	r := Rectangle{10,20}
	info(c)
	info(r)
}