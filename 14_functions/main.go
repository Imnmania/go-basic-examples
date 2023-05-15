package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is => ", result)

	result2 := proAdder(1, 2, 3, 4, 5, 6)
	fmt.Println("Pro result is => ", result2)

	proResInt, proResStr := proAdder2(1, 2, 3, 4, 5, 6)
	fmt.Println("Pro int result => ", proResInt)
	fmt.Println("Pro string result => ", proResStr)
}

func greeter() {
	fmt.Println("Hello from golang")
}

func greeterTwo() {
	fmt.Println("Hello from greeterTwo")
}

// --------------------------------------------
// function with two inputs and one return type
// --------------------------------------------
func adder(val1 int, val2 int) int {
	return val1 + val2
}

// --------------------
// function with vararg
// --------------------
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

// ----------------------------------
// function with multiple return type
// ----------------------------------
func proAdder2(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Returned text from proAdder2()"
}
