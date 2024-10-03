package main

import "fmt"

func sum(x, y int, s string) int {
	fmt.Println(s)
	return x + y
} //functio with multple return and quich returs

func calcuation(a, b int) (area, parameter int) {
	area = a * b
	parameter = a + b
	return
}

func main() {

	a, b := calcuation(3, 4)
	fmt.Println(a, b)

}
