package main

import (
	"fmt"
)

func VipEntry(guest map[string]bool, name string) {

	if _, ok := guest[name]; ok {
		fmt.Println("Welcome to the VIP Longue,", name)
	} else {
		fmt.Println("! Sorry,", name, "you are not on the guest list")
	}
}
func HonorRoll(grades map[string]int) {
	honorRoll := []string{}

	for name, score := range grades {
		if score > 80 {
			honorRoll = append(honorRoll, name)
		}
	}
	fmt.Println(honorRoll)
}
func UniqueName(names []string) {

	seen := map[string]bool{}
	namews := []string{}

	for _, name := range names {
		if !seen[name] {
			seen[name] = true
			namews = append(namews, name)
		}
	}

}
func main() {

	submittedNames := []string{"Alex", "Jordan", "Alex", "Sam", "Jordan", "Charlie"}
	UniqueName(submittedNames)

	//studentGrades := map[string]int{
	//	"Liam":  85,
	//	"Noah":  52,
	//	"Emma":  91,
	//	"Ava":   73,
	//	"Lucas": 45,
	//}
	//
	//HonorRoll(studentGrades)

	//guestList := map[string]bool{
	//	"Alice": true,
	//	"Bob":   true,
	//}

	//VipEntry(guestList, "ramesh")
	//VipEntry(guestList, "Bob")

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
