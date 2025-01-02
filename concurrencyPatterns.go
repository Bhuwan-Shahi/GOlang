package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func printNum(k int) {
	fmt.Println(k)
}

func main() {
	wg.Add(4)
	go printNum(3)
	go printNum(4)
	go printNum(5)
	go printNum(6)

	time.Sleep(time.Second * 1)
	wg.Done()
	fmt.Println("Hii")
}
