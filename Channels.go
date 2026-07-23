package main

import "fmt"

func add(a, b int, ch chan int) {
	ch <- a + b
}

func main() {
	ch := make(chan int)
	go add(2, 3, ch)
	x := <-ch
	fmt.Println(x)
}
