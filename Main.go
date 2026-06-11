package main

import "fmt"

//
//func myFunc(x, y int) (int, error) {
//	return x + y, nil
//}

func main() {
	//num, err := myFunc(10, 20)
	//fmt.Println(num, err)

	a := 2
	z := a > 1
	var age int
	fmt.Println("Please Enter your age:")
	_, err := fmt.Scan(&age)
	if err != nil {
		return
	}

	if age < 18 {
		fmt.Println("Are you dead or what!")
	} else {
		fmt.Println("Your seem eligible for the liscense")
	}
	if z == true {
		fmt.Println("Hello, World!")
	} else {
		fmt.Println("Hello, World! but in false")
	}

	fmt.Println(z)

	//Loops

	for i := 1; i < 10; i++ {
		fmt.Println(i, "Hello from :", i)
	}

	//While loop
	ab := 1
	for ab < 10 {
		fmt.Println("This is coming from while looping", ab)
		ab++
	}

	//Range Operator

	string := "Bhuwan Shahi"

	for _, char := range string {
		fmt.Printf("%c", char)

	}
	fmt.Println()

}
