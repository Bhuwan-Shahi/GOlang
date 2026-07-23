package main

import (
	"fmt"
	"time"
)

func add(a, b int, ch chan int, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
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

	go add(1, 2, ch1, 2)
	go add(3, 2, ch2, 2)

	for i := 0; i < 2; i++ {
		select {
		case x := <-ch1:
			fmt.Println(x)

		case y := <-ch2:
			fmt.Println(y)

		}
	}

	//Select allows us to handle which channel to use to handle the value

}
