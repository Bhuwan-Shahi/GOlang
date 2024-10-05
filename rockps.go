// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"strings"
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
// 	var playerChoice string
// 	choices := []string{"rock", "paper", "scissors"}
// 	for {
// 		fmt.Println("Enter your choice (rock, paper, scissors):")
// 		fmt.Scanln(&playerChoice)
// 		playerChoice = strings.ToLower(playerChoice)
// 		if Contains(choices, playerChoice) {
// 			fmt.Println("Your choice is:", playerChoice)
// 			break
// 		} else {
// 			fmt.Println("Wrong choice. Please enter again.")
// 		}
// 	}
// 	return playerChoice
// }

// func Contains(choices []string, element string) bool {
// 	for _, i := range choices {
// 		if i == element {
// 			return true
// 		}
// 	}
// 	return false
// }

// func deterMinewinner() {
// 	for playerWin < 5 && computerWin < 5 {
// 		computerChoice := getComputerChoice()
// 		playerChoice := getPlayerChoice()

// 		fmt.Println("Player chose:", playerChoice)
// 		fmt.Println("Computer chose:", computerChoice)

// 		if playerChoice == computerChoice {
// 			fmt.Println("__________________")
// 			fmt.Println("It's a tie")
// 			fmt.Println("__________________")
// 		} else {
// 			switch playerChoice {
// 			case "rock":
// 				if computerChoice == "scissors" {
// 					fmt.Println("__________________")
// 					fmt.Println("Player wins!!")
// 					fmt.Println("__________________")
// 					playerWin++
// 				} else {
// 					fmt.Println("__________________")
// 					fmt.Println("Computer Wins!!")
// 					fmt.Println("__________________")
// 					computerWin++
// 				}
// 			case "paper":
// 				if computerChoice == "rock" {
// 					fmt.Println("__________________")
// 					fmt.Println("Player wins!!")
// 					fmt.Println("__________________")
// 					playerWin++
// 				} else {
// 					fmt.Println("__________________")
// 					fmt.Println("Computer Wins!!")
// 					fmt.Println("__________________")
// 					computerWin++
// 				}
// 			case "scissors":
// 				if computerChoice == "paper" {
// 					fmt.Println("__________________")
// 					fmt.Println("Player wins!!")
// 					fmt.Println("__________________")
// 					playerWin++
// 				} else {
// 					fmt.Println("__________________")
// 					fmt.Println("Computer Wins!!")
// 					fmt.Println("__________________")
// 					computerWin++
// 				}
// 			}
// 		}

// 		fmt.Printf("Score - Player: %d, Computer: %d\n", playerWin, computerWin)
// 	}

// 	if playerWin == 5 {
// 		fmt.Println("Player wins the game!")
// 	} else {
// 		fmt.Println("Computer wins the game!")
// 	}
// }

// func main() {
// 	deterMinewinner()
// }
