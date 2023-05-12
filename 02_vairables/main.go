package main

import "fmt"

// in Go, const variables start with capital letter
const LoginToken string = "sdklfjslkdfj" // public

func main() {
	// string values
	var username string = "Niloy"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// bool values
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// int values
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// float values
	var smallFloat float32 = 255.45564567567567657567567
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit value
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style (walrus operator :=)
	// only allowed within functions
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)

	// constant variables
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
