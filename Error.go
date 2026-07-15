package main

import (
	"errors"
	"fmt"
)

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
	result, err := div(2, 2)
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		fmt.Println(result)
	}
}
