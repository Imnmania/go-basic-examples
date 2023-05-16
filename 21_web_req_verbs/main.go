package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web req in golang")

	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:1111/get"

	response, error := http.Get(myUrl)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	fmt.Println("Status Code => ", response.StatusCode)
	fmt.Println("Content length => ", response.ContentLength)

	/// way 1
	/* content, _ := io.ReadAll(response.Body)
	fmt.Println(content)
	fmt.Println(string(content)) */

	/// way 2
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is => ", byteCount) // same as length
	fmt.Println("The content is => ", responseString.String())
}
