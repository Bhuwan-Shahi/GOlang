package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://google.com"

func main() {
	fmt.Println("Handeling urls")

	fmt.Println(myUrl)

	//parsing the url
	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("%T", qparams)

}
