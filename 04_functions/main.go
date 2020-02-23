package main

import (
	"fmt"
)

func greeting(s string) string {
	return "hello, " + s
}

func getSum(a int, b int) int {
	return a + b
}

func main(){
	fmt.Println(greeting("hefei"))
	fmt.Println(getSum(1,2))

}