package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Hieght float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Hieght
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Println("The areas :is ", s.Area())
}
func main() {
	R := Rectangle{1, 2}
	C := Circle{2}
	PrintArea(R)
	PrintArea(C)

}
