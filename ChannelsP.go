package main

import (
	"errors"
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

var ErrInvalidAmoun = errors.New("Amount must be positive!")

type Account struct {
	Id      string
	Balance float64
}

func (acc *Account) deposit(amount float64) error {
	if acc == nil {
		return errors.New("Nil account pointer")
	}
	if amount <= 0 {
		return fmt.Errorf("deposit failed for account %s: %w", acc.Id, ErrInvalidAmoun)
	}
	acc.Balance += amount
	return nil

}

func main() {
	acc := Account{"Bhuwan", 0.0}

	fmt.Println(acc.deposit(12))
	fmt.Println(acc.Balance)
	//
	//ch := make(chan int)
	//arr := []int{1, 2, 3}
	//arr1 := []int{4, 5, 6}
	//go Ssum(arr, ch)
	//go Ssum(arr1, ch)
	//x := <-ch
	//y := <-ch
	//fmt.Println("the sum of the array is:")
	//fmt.Println(x + y)
	//x := <-ch
	//
	//fmt.Println(x)
}
