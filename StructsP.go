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

type Item struct {
	Name      string
	Quantity  int
	UnitPrice float64
}

func (I Item) TotalValue(inventory []Item) float64 {
	totalValue := 0.0

	for _, it := range inventory {
		totalValue += (float64(it.Quantity) * it.UnitPrice)
	}
	return totalValue
}

func main() {

	items := []Item{
		{"cup", 1, 2},
		{"cup", 1, 2},
		{"cup", 1, 2},
	}
	fmt.Println((Item{}).TotalValue(items))

	Cup := Product{"Paper Cup", 180, true}

	fmt.Println(Cup.IsAffordable(2))

}
