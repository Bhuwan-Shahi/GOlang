package main

import "fmt"

func deffered() {
	fmt.Println("This is from a defferd function")
	r := recover()
	fmt.Println(r)
}

func main() {
	fmt.Println("Hello World")
	defer deffered()
	panic("Hiii, I am causing the panice")
	fmt.Println("eh")
}
