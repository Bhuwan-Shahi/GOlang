package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrINsufficientFunds = errors.New("INsufficinet funds for this withdrawl")

func Withdraw(currentbalance float64, amount float64) (float64, error) {

	if amount > currentbalance {
		return currentbalance, ErrINsufficientFunds
	}
	return (currentbalance - amount), nil
}
func ParseAge(input string) (int, error) {
	integer, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalide age format : %w", err)
	}
	return integer, nil
}

func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by 0")

	}
	return a / b, nil
}

func deffered() {
	fmt.Println("This is from a defferd function")
	r := recover()
	fmt.Println(r)
}

func main() {
	fmt.Println(Withdraw(1200, 2100))
	fmt.Println(ParseAge("bhuwan"))
	result, err := div(2, 2)
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		fmt.Println(result)
	}
}
