package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Hello world.")

	circle := Circle{x:2, y:3, radius:1}
	rectangle := Rectangle{width:12, height:11}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))
}