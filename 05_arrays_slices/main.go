package main

import "fmt"

func main() {
	// Array
	var fruitArray [2]string
	// Assign values
	fruitArray[0] = "apple"
	fruitArray[1] = "orange"

	// Declare and assign
	fruitArray2 := []string{"Apple", "Orange", "grape"}
	fmt.Printf("%T\n", fruitArray)
	fmt.Printf("%T\n", fruitArray2)

	fmt.Println(fruitArray)
	fmt.Println(fruitArray2)
}