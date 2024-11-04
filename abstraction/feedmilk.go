package main

import "fmt"

type feedingMilk struct{}

func (w feedingMilk) feeds() {
	fmt.Println("The function is from feedingmilk.go and from funct name feeds ")
}
