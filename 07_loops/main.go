package main

import "fmt"

func main() {
	i:=1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i:=0; i < 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// fizzbuz
	for i:=10; i<100 ;i++ {
		if i % 3 == 0 {
			continue
		}
		if i == 50 {
			break
		}

		fmt.Printf("Number %d\n", i)
	}

}