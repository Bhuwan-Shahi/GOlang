package main

import (
	"sync"
)

type Jobs struct {
	id   int
	lock sync.Mutex
}

func (j Jobs) Runner(ch chan string) {
	x := "Your work is don of " + string(j.id)
	ch <- x
}
func main() {
	jobs := [20]Jobs{}
	for i := 0; i <= 20 ; i++ {
		jobs[i] =
 	}
}
