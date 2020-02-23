package main

import (
	"fmt"
)

func main() {
	var name = "brad"
	var age = 35
	const isCool =true
	var size = 1.7
	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}
