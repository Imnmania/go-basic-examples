package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	// fruitList[2] = "Bamboo"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is => ", fruitList)
	fmt.Println("Fruit list length is => ", len(fruitList))

	var vegList = [3]string{"potato", "bean", "mushroom"}
	fmt.Println("Veg list is => ", vegList)
	fmt.Println("Veg list length is => ", len(vegList))
}
