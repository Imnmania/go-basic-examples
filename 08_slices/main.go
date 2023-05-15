package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	// var fruitList = []string{}
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruit list => %T \n", fruitList)
	fmt.Println("Length of list => ", len(fruitList))
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Printf("Type of fruit list => %T \n", fruitList)
	fmt.Println("Length of list => ", len(fruitList))
	fmt.Println(fruitList)

	// fruitList = fruitList[1:]
	// fruitList = fruitList[1:3]
	fruitList = fruitList[:3]
	fmt.Println(fruitList)

	/// -----------
	/// use of make
	/// -----------
	highScores := make([]int, 4)
	highScores[0] = 235
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 //! causes error, since we gave it the size of 4
	fmt.Println(highScores)

	//? but adding works using "append", since it re-allocates the memory of the array
	highScores = append(highScores, 123, 345, 678)
	fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	/// -----------
	/// use of sort
	/// -----------
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	/// -----------------------------------------
	/// delete a value from slices based on index
	/// -----------------------------------------
	var courses = []string{"Flutter", "Java", "Kotlin", "Python", "C#"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
