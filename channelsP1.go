package main

import (
	"errors"
	"fmt"
)

var ErrOutofStock = errors.New("Out of stock right now")

type Iitem struct {
	name        string
	quantity    int
	disConitued bool
}

func (item *Iitem) Processor(name string, quantity int) error {
	if item == nil {
		return errors.New("nil item pointer")
	}
	if item.disConitued == true {
		return errors.New("This item is alredy discontinued!")
	}
	if item.quantity <= 0 {
		return fmt.Errorf("Stock for %s: %w", item.name, ErrOutofStock)
	}
	item.quantity -= quantity
	return nil
}
func Middle(iterm *Iitem, name string, quanitiy int, ch chan error) {
	ch <- iterm.Processor(name, quanitiy)
}
func main() {
	ch := make(chan error)
	i1 := Iitem{}
	i2 := Iitem{"snaks", 2, false}

	go Middle(&i1, "hello", 2, ch)
	go Middle(&i2, "snaks", 3, ch)
	go Middle(&i2, "snaks", 2, ch)
	err := <-ch
	err1 := <-ch
	err2 := <-ch

	if errors.Is(ErrOutofStock, err1) {
		fmt.Println(err1)
	}
	if errors.Is(ErrOutofStock, err) {
		fmt.Println(err1)
	}
	if errors.Is(ErrOutofStock, err2) {
		fmt.Println(err1)
	}
	fmt.Println("Remaining i1 quantity:", i1.quantity)
	fmt.Println("Remaining i2 quantity:", i2.quantity)
}
