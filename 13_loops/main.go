package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	/// ------
	/// Type 1
	/// ------
	/* for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	} */

	/// ------
	/// Type 2
	/// ------
	/* for i := range days {
		fmt.Println(days[i])
	} */

	/// ------
	/// Type 3
	/// ------
	/* for index, value := range days {
		fmt.Printf("index is %v and value is %v \n", index, value)
	} */
	/* for _, value := range days {
		fmt.Printf("value is %v \n", value)
	} */

	/// -----------------------------
	/// Type 3 (just like while loop)
	/// -----------------------------
	rougueValue := 1

	for rougueValue < 10 {

		/* if rougueValue == 5 {
			break
		} */
		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is => ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at the bottom!")
}
