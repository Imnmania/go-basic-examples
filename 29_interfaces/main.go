package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome to interfaces...")

	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, val := range shapes {
		printShapeInfo(val)
		fmt.Println("---")
	}
}

/*
	- In Golang, interfaces do not have any "implements" keyword.
	- They just work magically when the structs have the same (abstract) methods
	mentioned inside the interface. Just like hotswappable keys in mechanical
	keyboards, if they have the same shape and pins, they just work.
*/

// shape interface
type shape interface {
	area() float64
	circumf() float64
}

// square struct
type square struct {
	length float64
}

// circle struct
type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumferance of %T is: %0.2f \n", s, s.circumf())
}
