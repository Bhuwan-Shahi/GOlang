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
func getfollowerPrediction(follower_count int, inf_type string, months float64) float64 {
	factor := 2.00
	switch inf_type {
	case "fitness":
		factor = 4
	case "cosmetic":
		factor = 3
	}
	fmt.Println(math.Pow(factor, months))
	return float64(follower_count) * math.Pow(factor, months)

}
func getinfluencerScore(numberFollowers float64, average_Engagement_percentage float64) float64 {
	return average_Engagement_percentage * math.Log2(numberFollowers)
}
func factorial(ab int) int {
	if ab <= 1 {
		return ab
	}
	return ab * factorial(ab-1)
}
func main() {
	fmt.Println("FAct")
	fmt.Println(factorial(3))

	fmt.Println("\n", getinfluencerScore(40000, 0.3), "\n")

	fmt.Println("This is from the flollower prediction.")
	fmt.Println("\n", getfollowerPrediction(10, "fitness", 1), "\n")
	fmt.Println("The total spread count is:")
	fmt.Println(getEstimagedSpread([]float64{12, 12, 12}))

	fmt.Println(Min([]int{121, 2232, 32323, 2324, 5, 5, 6}))
	fmt.Println(Min([]int{}))

	fmt.Println(Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	fmt.Println(Sum([]int{2, 3, 2, 2, 2, 2, 2, 2, 2, 2}))

	fmt.Println(Average([]int{2, 3, 2, 2, 2, 2, 2, 2, 2, 2}))
}
