package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// greeting method (value receiver)
func (p Person)greet() string {
	return "hello my name is " + p.firstName + ", my age is " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person)hasBirthday() {
	p.age++	
}

func (p *Person)getMarried(spouseLastName string) {
	if p.gender == "male" {
		return 
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	persons := [2]Person{}
	persons[0] = Person{firstName:"Fei", lastName:"He", city:"hangzhou", gender:"male", age:32}
	persons[1] = Person{firstName:"Fei", lastName:"He", city:"hangzhou", gender:"male", age:32}
	fmt.Println(persons)

	fmt.Println(persons[0].greet())
	persons[0].hasBirthday()
	fmt.Println(persons[0].greet())

}
