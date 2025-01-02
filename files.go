// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func writeFile() {
// 	content := "This is going in file"

// 	file, err := os.Create("./Files_ok.txt")
// 	checkErr(err)

// 	length, err := io.WriteString(file, content)
// 	checkErr(err)
// 	fmt.Println("Length is :", length)
// 	defer file.Close()

// }

// func main() {
// 	fmt.Println("Welcome to file handeling")
// 	writeFile()
// 	readFile("./Files_ok.txt")

// }

// func readFile(filename string) {
// 	dataByte, err := os.ReadFile(filename)
// 	checkErr(err)

// 	fmt.Println("Text data inside is:", string(dataByte))
// }

// func checkErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
