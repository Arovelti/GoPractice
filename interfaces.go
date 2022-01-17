package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
	name   string
}

type Square struct {
	lenght float64
	name   string
}

type Triangle struct {
	base   float64
	height float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Circumference() float64 {
	return math.Pi * c.radius * 2
}

func (c Circle) String() string {
	return fmt.Sprintf(
		"Area is %v Circunference is %v",
		c.Area(), c.Circumference())
}

func (s Square) Area() float64 {
	return math.Pow(s.lenght, 2)
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func calculateArea(listOfShapes []Shape) {
	for _, shapes := range listOfShapes {
		fmt.Println("Area of shape is ", shapes.Area())
	}
}

func main() {
	c1 := Circle{radius: 5, name: "c1"}
	s1 := Square{lenght: 4, name: "s1"}
	t1 := Triangle{base: 15, height: 20}
	shapes := []Shape{c1, s1, t1}
	fmt.Println(shapes)
	fmt.Println(c1)
}
