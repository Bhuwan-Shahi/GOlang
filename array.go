package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello ,world")

	// grades := []int{2, 3, 4, 5, 3, 5, 5}
	// fmt.Println(grades)
	// fmt.Println(len(grades))
	// fmt.Println(cap(grades))
	// fmt.Println(grades[4:7])

	a := make([]string, 0, 342)
	a = append(a, "hello")
	fmt.Println(a)
	fmt.Println(cap(a))
	b := []int{2, 3, 4, 5}
	c := b

	fmt.Println(c)

	d := 5

	if d == 5 {
		fmt.Println("hello world ")
	}
}
