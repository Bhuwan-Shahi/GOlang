package main

import (
	"fmt"
	"math"
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

func getEstimagedSpread(audiencesFlowers []float64) float64 {
	var numofFllowers float64
	numofFllowers = float64(len(audiencesFlowers))

	if numofFllowers == 0 {
		return 0
	}
	sum := 0.00
	for _, count := range audiencesFlowers {
		sum += count
	}
	sum /= numofFllowers
	bracket := math.Pow(float64(numofFllowers), 1.2)
	return float64(float64(sum) * bracket)
}

func main() {
	fmt.Println("The total spread count is:")
	fmt.Println(getEstimagedSpread([]float64{12, 12, 12}))
	fmt.Println(Min([]int{121, 2232, 32323, 2324, 5, 5, 6}))
	fmt.Println(Min([]int{}))

	fmt.Println(Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	fmt.Println(Sum([]int{2, 3, 2, 2, 2, 2, 2, 2, 2, 2}))

	fmt.Println(Average([]int{2, 3, 2, 2, 2, 2, 2, 2, 2, 2}))
}
