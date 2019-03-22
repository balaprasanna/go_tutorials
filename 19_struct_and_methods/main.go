package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func main() {
	fmt.Println("")
	c1 := Circle{radius: 5.0}
	c2 := Circle{9.0}
	fmt.Println("Area = ", c1.Area())
	fmt.Println("Area = ", c2.Area())

	s1 := Square{side: 10}
	fmt.Println("Area = ", s1.Area())
}
