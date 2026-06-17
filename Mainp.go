package main

import (
	"fmt"
)

func main() {

	submittedNames := []string{"Alex", "Jordan", "Alex", "Sam", "Jordan", "Charlie"}
	uniqueName(submittedNames)
	//inventory := map[string]int{

	//	"Milk":   5,
	//	"Bread":  0,
	//	"Eggs":   12,
	//	"Butter": 0,
	//}
	//
	//OutofStock(inventory)
	//var numbers []int
	//randomNumbers := []int{34, 21, 65, 56, 78, 90, 23, 12}

	//Frequency := []string{"ram", "apple", "ball", "appli", "ball", "ball", "ball", "ball", "apple"}

	//fmt.Println(WordFrequencyCounter(Frequency))
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
	//SumofSlice(randomNumbers)
	//	statusCodes := map[string]int{
	//		"en-US": 200,
	//		"fr-FR": 404,
	//	}
	//
	//	CommaOk(statusCodes, "am")
	//	CommaOk(statusCodes, "en-US")
}
func uniqueName() {

}
func OutofStock(a map[string]int) {
	outstockslice := []string{}

	for i, value := range a {
		if value == 0 {
			//fmt.Println(value)
			outstockslice = append(outstockslice, i)
		}

	}
	fmt.Println(outstockslice)
}

//func WordFrequencyCounter(ram []string) map[string]int {
//	analyzed := map[string]int{}
//
//	for _, value := range ram {
//		analyzed[value]++
//	}
//
//	return analyzed
//}

func CommaOk(ram map[string]int, a string) {

	value, ok := ram[a]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("sorry")
	}
}

//func SumofSlice(Slice []int) {
//	sum := 0
//	for _, value := range Slice {
//		sum += value
//	}
//	fmt.Println(sum)
//}
