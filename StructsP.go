package main

import (
	"fmt"
)

type Product struct {
	Name    string
	Price   float64
	InStock bool
}

func (P Product) IsAffordable(budget float64) bool {
	return P.InStock && P.Price <= budget
}

func main() {
	Cup := Product{"Paper Cup", 180, true}

	fmt.Println(Cup.IsAffordable(2))

}
