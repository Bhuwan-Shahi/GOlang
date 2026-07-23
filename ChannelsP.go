package main

import (
	"fmt"
	"time"
)

func Pong(ch chan string, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	ch <- "Pong"
}

func Ssum(arr []int, ch chan int) {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	ch <- sum
}

func main() {

	ch := make(chan int)
	arr := []int{1, 2, 3}
	arr1 := []int{4, 5, 6}
	go Ssum(arr, ch)
	go Ssum(arr1, ch)
	x := <-ch
	y := <-ch
	fmt.Println("the sum of the array is:")
	fmt.Println(x + y)
	//x := <-ch
	//
	//fmt.Println(x)
}
