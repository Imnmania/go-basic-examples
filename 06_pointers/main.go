package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is => ", ptr)

	myNumber := 23

	/// Reference => &
	var ptr = &myNumber
	fmt.Println("Value of actual pointer => ", ptr)
	/// Value of Reference => *
	fmt.Println("Value of actual pointer => ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is => ", *ptr)
	fmt.Println("Old value is => ", myNumber)
}
