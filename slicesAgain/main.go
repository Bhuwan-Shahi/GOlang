package main

import "fmt"

var users = []string{}
var nums = []int{}

func addUsers(user ...string) {
	users = append(users, user...)
}

func main() {
	addUsers("bhuwan", "ram", "sita")
	fmt.Println(users)
	addNumbers(1, 3, 3, 4, 2, 55, 5, 2, 43, 3, 2)
	fmt.Println(nums)

}

func addNumbers(num ...int) {
	nums = append(nums, num...)
}
