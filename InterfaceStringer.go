package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type String interface {
	String() string
}

func (p Person) String() string {
	return fmt.Sprintf("%v %v (%d year old)",
		p.FirstName, p.LastName, p.Age)
}

func main() {
	me := Person{"Dima", "Mikhed", 29}
	fmt.Println(me)
}
