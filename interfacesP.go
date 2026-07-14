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

type PaymentProcessor interface {
	Pay(amount float64) bool
}
type PayPal struct {
	Email string
}

type CreditCard struct {
	CardNumber string
}

func (p PayPal) Pay(amount float64) bool {
	fmt.Println("Paid $", amount, "using PayaPal Account:", p.Email)
	return true
}
func (c CreditCard) Pay(amount float64) bool {
	lastfour := "xxxx"
	if len(c.CardNumber) >= 4 {
		lastfour = c.CardNumber[len(c.CardNumber)-4:]
	}
	fmt.Println("Charged $", amount, "to card ending in: ", lastfour)
	return true
}

func Checkout(p PaymentProcessor, total float64) {
	p.Pay(total)
}

func main() {

	p := PayPal{"Ramprasd@gmail.com"}
	c := CreditCard{"1233345"}

	Checkout(p, 123.2)
	Checkout(c, 123.2)

	R := Rectangle{1, 2}
	C := Circle{2}
	PrintArea(R)
	PrintArea(C)
}
