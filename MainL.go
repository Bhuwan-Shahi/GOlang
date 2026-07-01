package main

import (
	"fmt"
)

//
//func myFunc(x, y int) (int, error) {
//	return x + y, nil
//}

type Employee struct {
	name     string
	position string
}

func (E Employee) getname() string {
	fmt.Println("This is coming from the function inside struct")
	return E.name
}

func (e Employee) getPosition() string {
	fmt.Println("This is the get Position Function")
	return e.position

}
func main() {

	E1 := Employee{"Himal", "Manager"}
	fmt.Println(E1.getname())
	fmt.Println(E1.getPosition())
	//fmt.Println(E1.position)
	//fmt.Println(E1.name)

	//fmt.Println(E1)

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

	stringe := "Bhuwan Shahi"

	for _, char := range stringe {
		fmt.Printf("%c", char)

	}
	fmt.Println()

	//Arrays

	array := [...][2]int{{1, 2}, {2, 3}, {3, 4}}
	//To change the index of aray you have to do is initialize new
	array[0] = [2]int{34, 45}
	fmt.Println(array)

	for i, value := range array {
		fmt.Println(i, value)
	}
	//We cannot mutate the original array

	//slices

	//Pointer -> to the start
	//Lengh -> of the slice
	//Capacity -> how much more can this be expanded

	arr := [5]int{1, 2, 3, 4, 5}

	sl1 := arr[1:5]
	println("This is for slices only")
	fmt.Println(sl1, len(sl1), cap(sl1))

	fmt.Println("This is slices without and array")

	sl2 := []int{}
	fmt.Println(sl2, len(sl2), cap(sl2))

	for _, value := range sl1 {
		fmt.Println(value)
	}
	//slice get mutated while passing to or from function

	mp := map[string]int{"ram": 12, "shyam": 13}
	fmt.Println(mp)

}
