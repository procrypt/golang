package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Square struct {
	length float64
	width float64
}

func (c Circle) Area() float64 {
	return 3.14 * math.Sqrt(c.radius)
}

func (s Square) Area() float64 {
	return s.length*s.width
}

type Shape interface {
	Area() float64
}

func info(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	sq := Square{
		length: 10,
		width: 20,
	}

	ci := Circle{
		radius: 50,
	}

	info(sq)
	info(ci)
}