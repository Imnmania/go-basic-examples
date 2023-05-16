package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=paypal"

func main() {
	fmt.Println("Welcome to urls in golang")
	fmt.Println(myUrl)

	// parsing
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The type of query params are => %T \n", qParams)

	fmt.Println(qParams["coursename"])
	fmt.Println(qParams["payment"])

	for _, val := range qParams {
		fmt.Println("Param is => ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=niloy",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
