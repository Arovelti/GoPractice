package main

import (
	"fmt"
	"strconv"
)

func convertIntToString() {
	var input string
	fmt.Println("Enter your age: ")
	fmt.Scanf("%s", &input)

	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your age is:", age)
	}
}

func convToBool() {
	b, err := strconv.ParseBool("t")
	fmt.Println(b)
	fmt.Println(err)
	fmt.Printf("%T\n", b)
}

func main() {
	convertIntToString()
	convToBool()

	num1 := 5
	num2 := int64(num1)
	fmt.Printf("%T\n", num2)
}
