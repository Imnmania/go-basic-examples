package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in golang...")

	/* var myNumberOne int = 2
	var myNumberTwo float64 = 4
	fmt.Println("The sum => ", myNumberOne+int(myNumberTwo)) */

	// ---------------------
	// Random Numbers (math)
	// ---------------------
	/* r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	fmt.Println(r.Intn(5) + 1) */

	// -----------------------
	// Random Numbers (crypto)
	// -----------------------
	randomNum, _ := rand.Int(rand.Reader, big.NewInt(6))
	fmt.Println(randomNum)
}
