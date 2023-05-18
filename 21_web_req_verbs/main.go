package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web req in golang")

	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

// -----------
// Get Request
// -----------
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

// ---------------------------
// Post Request with Json Data
// ---------------------------
func PerformPostRequest() {
	const myUrl = "http://localhost:1111/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": 0,
			"platform": "learncodeonline.in"
		}
	`)

	contentType := "application/json"

	response, error := http.Post(myUrl, contentType, requestBody)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

// ---------------------------
// Post Request with Form Body
// ---------------------------
func PerformPostFormRequest() {
	const myUrl = "http://localhost:1111/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "Niloy")
	data.Add("lastname", "Biswas")
	data.Add("email", "n@b.com")

	response, error := http.PostForm(myUrl, data)

	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
