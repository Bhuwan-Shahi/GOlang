package main

import "fmt"

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

func main() {
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
