package main

import (
	"fmt"
	"math/rand"
	"time"
)

//DISPLAY THE BOARD
//GENERATE A RANDOM CHOICE OF COMPUTER
//FILL THA BOARD
//TAKE INPUT FROM THE USER
//FILL THE BOARD
//WRITE THE WINNING LOGIC

var Board = [][]string{
	{"_", "_", "_"},
	{"_", "_", "_"},
	{"_", "_", "_"},
}
var playerChoice string
var win bool

func displayBoard() {
	for _, j := range Board {
		fmt.Println(j)
	}
}

func playerInput() {
	fmt.Println("Choose between(O,X):")
	fmt.Scanln(&playerChoice)
}

func playerPlaces(playerChoice string) {
	var x, y int
	fmt.Println("Enter your horizontal choice:(0,1,2): ")
	fmt.Scanln(&x)
	fmt.Println("Enter your vertical choice:(0,1,2): ")
	fmt.Scanln(&y)
	Board[x][y] = playerChoice

}

func computerChoiceGeneration() {
	var ComputerChoice string
	if playerChoice == "x" {
		computerChoice = "o"
	} else {
		computerChoice = "x"
	}
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	x := generator.Intn(2)
	y := generator.Intn(2)

	Board[x][y] = computerChoice

}

func winCheck(choice string) bool {
	for i := 0; i < len(Board); i++ {
		if Board[i][0] == choice && Board[i][1] == choice && Board[i][2] == choice {
			return true
		}
	}
	for j := 0; j < len(Board); j++ {
		if Board[0][j] == choice && Board[1][j] == choice && Board[2][j] == choice {
			return true
		}
	}
	if Board[0][0] == choice && Board[1][1] == choice && Board[2][2] == choice {
		return true
	}
	if Board[0][2] == choice && Board[1][1] == choice && Board[2][0] == choice {
		return true
	}
	return false

}

func over() {
	playerInput()
	for !win {
		displayBoard()

		playerPlaces(playerChoice)
		computerChoiceGeneration()
		if winCheck(playerChoice) {
			fmt.Println("Player wins!!")
			win = true
			displayBoard()
			break
		}el

	}
}

func main() {
	fmt.Println("This is a tic tac toe ")
	over()

}
