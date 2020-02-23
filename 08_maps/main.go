package main

import "fmt"

func main() {
	//emails := make(map[string]string)
	emails := map[string]string{"mike":"mike@gmail.com"}
	emails["hefei"] = "hefei.lark@bytedance.com"
	emails["zjy"] = "zjy-111@qq.com"
	fmt.Println(emails)
	delete(emails, "mike")
	fmt.Println(emails)
}