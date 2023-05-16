package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go inside a file - niloybiswas.me"
	file, err := os.Create("./mygofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is => ", length)
	defer file.Close()
	readFile("./mygofile.txt")
}

func readFile(filename string) {
	// dataByte, err := ioutil.ReadFile(filename) //! ReadFile got deprecated on ioutil package
	dataByte, err := os.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n ", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
