package main

import "fmt"

func add(a, b int, ch chan int) {
	fmt.Println(a + b)
	ch <- a + b
}

func main() {
	ch := make(chan int)
	go add(2, 3, ch)
	x := <-ch

	go add(6, 3, ch)
	x = <-ch
	go add(5, 3, ch)
	x = <-ch
	go add(4, 3, ch)
	x = <-ch
	go add(3, 3, ch)
	x = <-ch
	fmt.Println(x)
}
