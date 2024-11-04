package main

import "fmt"

type distinctisgone struct{}

func (d distinctisgone) distinct() {
	fmt.Println("This is from the distinct")
}
