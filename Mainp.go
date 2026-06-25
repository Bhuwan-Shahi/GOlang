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
func UpdateShipment(currentStock map[string]int, newShipment []string) {

	for _, value := range newShipment {
		currentStock[value] += 1
	}
	fmt.Println(currentStock)
}
func CalculateAverage(prices map[string]int) float64 {
	var average float64

	for i, _ := range prices {
		average += float64(prices[i])
	}
	average = average / float64(len(prices))

	return average
}

func GroupByLetters(names []string) {
	group := map[string][]string{}
	for _, name := range names {
		firstLetter := string(name[0])
		group[firstLetter] = append(group[firstLetter], name)
	}
	fmt.Println(group)
}

func MultiCalculator(a, b int) (sum, sub, mul int, div float64) {

	sum = a + b
	sub = a - b
	mul = a * b

	if b == 0 {
		div = 0.0
		return
	}
	div = float64(a) / float64(b)
	return sum, sub, mul, div
}
func main() {
	fmt.Println(MultiCalculator(4, 4))
	//names := []string{"Alice", "Bob", "Charlie", "Anna", "Bill"}
	//
	//GroupByLetters(names)
	//menuPrices := map[string]int{
	//	"Burger": 15,
	//	"Fries":  5,
	//	"Soda":   3,
	//	"Shake":  7,
	//}
	//
	//fmt.Println(CalculateAverage(menuPrices))

	//currentStock := map[string]int{"shoes": 5, "hats": 2}
	//newShipment := []string{"shoes", "shirts", "shoes", "hats"}
	//
	//UpdateShipment(currentStock, newShipment)

	//submittedNames := []string{"Alex", "Jordan", "Alex", "Sam", "Jordan", "Charlie"}
	//UniqueName(submittedNames)

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
