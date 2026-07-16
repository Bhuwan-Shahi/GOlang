package main

import (
	"fmt"
)

//1. Finding Minimum

func Min(number []int) int {
	if len(number) == 0 {
		return 0
	}
	min := number[0]

	for _, i := range number {
		if i < min {
			min = i
		}
	}
	return min
}
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
func Average(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum / len(numbers)
}

func main() {
	fmt.Println(Min([]int{121, 2232, 32323, 2324, 5, 5, 6}))
	fmt.Println(Min([]int{}))

	fmt.Println(Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	fmt.Println(Sum([]int{2, 3, 2, 2, 2, 2, 2, 2, 2, 2}))

	fmt.Println(Average([]int{2, 3, 2, 2, 2, 2, 2, 2, 2, 2}))
}
