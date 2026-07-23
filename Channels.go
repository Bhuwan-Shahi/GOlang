package main

import (
	"fmt"
	"time"
)

func add(a, b int, ch chan int, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println(a + b)
	ch <- a + b
}

func main() {
	//ch := make(chan int)
	//go add(2, 3, ch)
	//x := <-ch
	//
	//go add(6, 3, ch)
	//x = <-ch
	//go add(5, 3, ch)
	//x = <-ch
	//go add(4, 3, ch)
	//x = <-ch
	//go add(3, 3, ch)
	//x = <-ch
	//fmt.Println(x)
	//

	fmt.Println("Adding Mulitple Channels")
	ch1 := make(chan int)
	ch2 := make(chan int)

	go add(1, 2, ch1, 1)
	go add(3, 2, ch2, 2)

	x := <-ch2
	y := <-ch1

	fmt.Println(x, y)
}
