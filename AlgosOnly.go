package main

import "fmt"

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

func main() {
	fmt.Println(Min([]int{121, 2232, 32323, 2324, 5, 5, 6}))
	fmt.Println(Min([]int{}))

}
