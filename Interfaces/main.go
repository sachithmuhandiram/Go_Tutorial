package main

import (
	"fmt"
	"math"
)

type Circle struct {
	shape_type string
	radius     float64
}

type Square struct {
	shape_type string
	side       float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// area() receiving square
func (s Square) area() float64 {
	return s.side * s.side
}

/*
	Both functions have the same name area(),
	but golang consider them as two functions as they receive
	different structs
*/

type Shapes interface {
	area() float64
}

/*
	This Shapes interface has area() method.
	Anything access to area() is considered as a Shape.
	Both Circle and Square strcuts can access to it
*/
func info(sh Shapes) {
	fmt.Println("From Shape interface : ")
	fmt.Println(sh)
	fmt.Println(sh.area())
}

func main() {
	c1 := Circle{
		shape_type: "circle",
		radius:     3.5,
	}

	fmt.Println("Area of the circle is : ", c1.area())

	info(c1)

	s1 := Square{
		shape_type: "square",
		side:       10.5,
	}
	// this struct has only one variable, no need to parameter wise addition
	fmt.Println("Area of square : ", s1.area())

	info(s1)
}
