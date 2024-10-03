// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// var playerWin, computerWin int

// func getComputerChoice() string {
// 	source := rand.NewSource(time.Now().UnixNano())

// 	generator := rand.New(source)

// 	choices := []string{"rock", "paper", "scissors"}
// 	randomIndex := generator.Intn(len(choices))
// 	randomChoice := choices[randomIndex]
// 	return randomChoice

// }
// func getPlayerChoice() string {
// 	var playerchoice string

// 	fmt.Println("Enter your choice(rock, paper, scissors):")
// 	fmt.Scanln(&playerchoice)
// 	fmt.Println("Hello,", playerchoice)
// 	return playerchoice
// }

// func main() {
// 	fmt.Println(getComputerChoice())
// 	fmt.Println(getPlayerChoice())
// }