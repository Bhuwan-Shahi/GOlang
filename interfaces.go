package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	height, breadth float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.breadth * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func calcArea(s Shape) float64 {
	return s.Area()
}

func main() {
	fmt.Println("Hello world")
	rect := Rectangle{
		height:  34,
		breadth: 1,
	}
	c := Circle{
		radius: 2,
	}

	fmt.Println("The area of rectangle is:", calcArea(rect))
	fmt.Println("The area of circle is:", calcArea(c))

}
