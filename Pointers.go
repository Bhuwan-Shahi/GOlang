package main

import (
	"errors"
	"fmt"
)

type Book struct {
	id    int
	title string
}

func (b *Book) setTitle(title string) {
	b.title = title
}

func change(x *int) {
	*x = 100
	fmt.Println(*x)
}

func Swap(x, y *int) {
	*x, *y = *y, *x

}

type Wallet struct {
	Balance float64
}

func (w *Wallet) Deposit(amount float64) {
	w.Balance += amount
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.Balance {
		return errors.New("Amount greather than balance")
	}
	w.Balance -= amount
	return nil
}
func (w *Wallet) PrintWallet() {
	if w == nil {
		fmt.Println("Error: No Wallet Found")
		return
	}
	fmt.Println("||||||||||")
	fmt.Println("The current balance is:", w.Balance)
	fmt.Println("||||||||||")

}

func main() {
	var voidWallet *Wallet
	voidWallet.PrintWallet()

	bhuwansWallet := Wallet{100.00}

	bhuwansWallet.Deposit(50.00)
	fmt.Println(bhuwansWallet.Withdraw(30))
	fmt.Println(bhuwansWallet)
	bhuwansWallet.PrintWallet()

	b := Book{1, "Bhuwan ko Book"}
	b.setTitle("Bhuwan ko book aaba ram ko vayo because pointer ko refrence")
	fmt.Println(b)
	x := 0
	fmt.Println(x)
	y := &x
	z := &y

	*y = 8
	fmt.Println(**z)

	fmt.Println(x)
	yy := 1

	change(&yy)
	println(yy)

	fmt.Println("The output is for pointer pracitce quesiton only")
	yyy := 12
	zzz := 13
	fmt.Println(yyy, zzz)
	Swap(&yyy, &zzz)
	fmt.Println(yyy, zzz)

}
