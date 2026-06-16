package main

import "fmt"

func main() {
	//var numbers []int
	randomNumbers := []int{34, 21, 65, 56, 78, 90, 23, 12}
	//largestNumber := randomNumbers[0]
	//
	//for _, value := range randomNumbers {
	//	if value > largestNumber {
	//		largestNumber = value
	//
	//	}
	//}
	//fmt.Println("_______")
	//fmt.Println(largestNumber)

	//fmt.Println("Now creating array")
	//
	//for i := 1; i <= 10; i++ {
	//	numbers = append(numbers, i)
	//
	//}
	//
	//fmt.Println(numbers)
	//
	//fmt.Println("Now printing the even number only")
	//for _, value := range numbers {
	//	if value%2 == 0 {
	//		fmt.Println(value)
	//	}
	//}
	SumofSlice(randomNumbers)
}

func SumofSlice(Slice []int) {
	sum := 0
	for _, value := range Slice {
		sum += value
	}
	fmt.Println(sum)
}
