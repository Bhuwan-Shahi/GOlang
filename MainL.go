package main

import (
	"fmt"
)

//
//func myFunc(x, y int) (int, error) {
//	return x + y, nil
//}

type Hobby struct {
	Name     string
	Duration int
}

type Employee struct {
	name     string
	position string
	hobbies  Hobby
}

func (E Employee) getname() string {
	fmt.Println("This is coming from the function inside struct")
	return E.name
}

func (e Employee) getPosition() string {
	fmt.Println("This is the get Position Function")
	return e.position

}
func (e Employee) getHobbies() Hobby {
	return e.hobbies
}

type Shape interface {
	getPerimter() int
	getArea() int
}

type triangle struct {
	length  int
	breadth int
	height  int
}
type square struct {
	length int
}

func (s square) getPerimter() int {
	return 4 * s.length
}

func (s square) getArea() int {
	return 4 * s.length
}

func (t triangle) getPerimter() int {
	return t.length + t.breadth + t.height
}

func (t triangle) getArea() int {
	return t.length * t.breadth
}
func main() {
	var shapes []Shape = []Shape{triangle{1, 2, 3}, square{1}}
	perimeter := 0

	for _, shape := range shapes {
		perimeter += shape.getArea()
	}
	fmt.Println("This is sum from the array:", perimeter)

	fmt.Println("Array of the interfaces")
	var s Shape = triangle{1, 2, 4}
	fmt.Println(s.getPerimter())
	fmt.Println(s.getArea())
	fmt.Println("this is where interface ends")

	E1 := Employee{"Himal", "Manager", Hobby{"Reading", 5}}
	fmt.Println(E1.getname())
	fmt.Println(E1.getPosition())
	fmt.Println(E1.getHobbies())
	fmt.Println(E1)
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
