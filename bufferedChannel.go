package main

import (
	"fmt"
	"time"
)

func work(ch <-chan string) {
	for i := 0; i < 4; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
}

func Worker(ch <-chan int) {
	for i := 0; i < 3; i++ {
		task := <-ch
		fmt.Println(task)
	}
}

func main() {

	//ch := make(chan string, 3)
	tasks := make(chan int, 2)
	go Worker(tasks)
	for i := 0; i <= 2; i++ {
		tasks <- i * 100
	}

	time.Sleep(100 * time.Millisecond)

	//go work(ch)
	//for i := 0; i <= 4; i++ {
	//	ch <- fmt.Sprint("log%d", i)
	//	fmt.Printf("Sent log%d to channel\n", i)
	//}
	//time.Sleep(100 * time.Millisecond)

}
