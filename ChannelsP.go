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
func Middleman(acc *Account, amount float64, ch chan error) {
	ch <- acc.deposit(amount)
}

func main() {
	ch := make(chan error)
	acc := Account{"Bhuwan", 0.0}

	go Middleman(&acc, 12, ch)
	go Middleman(&acc, 0, ch)
	err1 := <-ch
	err2 := <-ch

	if errors.Is(err1, ErrInvalidAmoun) {
		fmt.Println("Caught error:", err1)
	}
	if errors.Is(err2, ErrInvalidAmoun) {
		fmt.Println("Caught error:", err2)
	}

	fmt.Println("Final Balance:", acc.Balance)

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
