package main

import (
	"fmt"
)

func aadd(x, y int, ch chan int) {
	ch <- x + y
}

func main() {
	ch := make(chan int)
	go aadd(4, 5, ch)
	x := <-ch
	fmt.Println(x)
}
