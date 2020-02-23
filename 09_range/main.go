package main

import "fmt"

func main() {

	ids := []int{1,2,3,345,23,52,4,234,2,4,23,4}

	emails := map[string]string{"hefei":"hefei@qq.com"}

	for i, id := range ids{
		fmt.Printf("%d is %d \n", i, id)
	}

	for k,v :=range emails {
		fmt.Printf("%s is %s \n", k, v)
	}


}