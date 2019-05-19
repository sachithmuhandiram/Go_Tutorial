package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shapes interface {
	area() float64
}

/*
	This Shapes interface has area() method.
	Anything access to area() is considered as a Shape.
*/
func info(sh Shapes) {
	fmt.Println("From Shape interface : ")
	fmt.Println(sh)
	fmt.Println(sh.area())
}

func main() {
	c1 := Circle{
		radius: 3.5,
	}

	fmt.Println("Area of the circle is : ", c1.area())

	info(c1)

}
