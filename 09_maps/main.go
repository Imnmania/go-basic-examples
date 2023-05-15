package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps")

	/// -----------
	/// declare map
	/// -----------
	languages := make(map[string]string)

	/// ----------
	/// add to map
	/// ----------
	languages["js"] = "Javascript"
	languages["ts"] = "Typescript"
	languages["py"] = "Python"

	fmt.Println("List of all languages => ", languages)
	fmt.Println("js shorts for => ", languages["js"])

	/// ---------------
	/// delete from map
	/// ---------------
	delete(languages, "js")
	fmt.Println("List of all languages => ", languages)

	/// ------------
	/// loops in map
	/// ------------
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
