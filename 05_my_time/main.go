package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// "01-02-2006 15:04:05 Monday" is from official documentation
	// must use these exact times/dates in order to format date time in golang.
	// no "DD-MM-YYYY HH:mm:ss A" like any other programming languages
	fmt.Println(presentTime.Format("02-01-06 15:04:05 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 03:04 PM Monday"))

	// create our own time
	createdDate := time.Date(2020, time.July, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-06 Monday"))
}
