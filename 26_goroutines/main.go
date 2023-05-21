package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Welcome to goroutine...")
	/* go greeter("Hello")
	greeter("World") */

	websiteList := []string{
		"https://google.com",
		"https://pub.dev",
		"https://go.dev",
		"https://github.com",
		"https://nyaa.si",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

/* func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
} */

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOps in endpoint")
		return
	}
	fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
}
