package main

import "fmt"

type Greeting struct {
	name   string
	rollno int
}

func (g *Greeting) greets() {
	fmt.Println("Hello", g.name, "your rollno is :", g.rollno)
	g.name = "not existent"
}

func main() {
	g := Greeting{
		name:   "Bhuwan",
		rollno: 8,
	}
	fmt.Println("The original name is :", g.name)
	g.greets()
	fmt.Println(g.name)

}

//THESE TYPE OF FUNCITON ONLY GETS THE COPY OF THE ACUTAL STRUCT.
//IF WE CHANGE THE STRUCT VALUE IN THE METHOD THE REAL DATA WONT BE CHANGE
//IN ORDER TO PERFORM THAT WE HAVE TO MAKE CHANGES IN IN THE PASSING OF THE STRUCT AND GIVE IT WITH TH POIONTER VARIBLAE
