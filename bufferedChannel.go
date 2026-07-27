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

func main() {
	ch := make(chan string, 3)

	go work(ch)
	for i := 0; i <= 4; i++ {
		ch <- fmt.Sprint("log%d", i)
		fmt.Printf("Sent log%d to channel\n", i)
	}
	time.Sleep(100 * time.Millisecond)
}
