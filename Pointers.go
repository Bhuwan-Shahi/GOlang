package main

import "fmt"

func main() {
	x := 0
	fmt.Println(x)
	y := &x

	*y = 8

	fmt.Println(x)

}
