package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web request in golang")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type => %T \n", res)

	defer res.Body.Close() // callers responsibility to close the connection

	// dataByte, err := ioutil.ReadAll(res.Body) //! deprecated
	dataByte, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataByte)
	fmt.Println(content)
}
