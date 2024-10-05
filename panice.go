// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("helo wev dev!"))
// 	})
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	a := func() {
// 		fmt.Println("hello world for the anonymous function")
// 	}
// 	fmt.Print(a)

// }
