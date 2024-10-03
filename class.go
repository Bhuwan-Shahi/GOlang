package main

import (
	"fmt"
	"reflect"
)

type College struct {
	name       string
	affliation string
	levels     []int
}
type car struct {
	name   string
	model  []int
	runsIn string
}
type hello struct{
	hello stirng;

}

func compare(a, b College) string {
	if reflect.DeepEqual(a, b) {
		fmt.Print("they are equal")
		return "yes"
	}
	return "no"
}

func main() {
	padmashree := College{
		name:       "nist",
		affliation: "tu",
		levels:     []int{3, 3, 4, 5, 6, 6},
	}
	padmashree1 := College{
		name:       "nist",
		affliation: "tu",
		levels:     []int{3, 3, 4, 5, 6, 6},
	}
	fmt.Println(padmashree)

	car1 := car{
		na me:   "buggati ciron",
		model:  []int{2013, 2014, 2015, 2016},
		runsIn: "disel",
	}
	fmt.Println(car1)
	fmt.Println(car1.model[2])

	k := compare(padmashree, padmashree1)
	fmt.Println()
	fmt.Println(k)

	car2 := car {}
	car2.
	
	fmt.Println(car2.name)



}
