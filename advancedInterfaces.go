package main

import "fmt"

type Mammals interface {
	feedMilk() string
}

type Birds interface {
	laysEggs() string
}

type Bat interface {
	Mammals
	Birds
}

type Animals struct {
	kingdom string
	class   string
}

type animalError struct {
	msg string
}

func (a Animals) feedMilk() string {
	if a.class != "bird" {
		return "Yes the feed milk"
	} else {
		return "no they dont feed milk"
	}
}

func (a Animals) laysEggs() string {
	if a.class != "mammals" {
		return "They lay eggs"
	}
	return "Thy don't lay eggs"
}

func check(a Animals) {
	fmt.Println(a.feedMilk())
	fmt.Println(a.laysEggs())
}

func main() {

	pigeon := Animals{
		kingdom: "animal",
		class:   "bird",
	}
	hippo := Animals{
		kingdom: "animal",
		class:   "mammals",
	}

	check(pigeon)
	check(hippo)

}
