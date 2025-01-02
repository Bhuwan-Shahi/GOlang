// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// const url = "https://google.com"

// func main() {
// 	fmt.Println("web request")

// 	response, err := http.Get(url)

// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Println("Response i :%T", response)

// 	defer response.Body.Close() //callers responsibility to close the request

// 	databytes, err := io.ReadAll(response.Body)

// 	if err != nil {
// 		panic(err)
// 	}
// 	content := string(databytes)

// 	fmt.Println(content)
// }
