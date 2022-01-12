package main

import "fmt"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	x := 15
	y := 20
	minNum := Min(x, y)
	maxNum := Max(x, y)
	fmt.Printf("Mix from %v and %v is: %v\n", x, y, minNum)
	fmt.Printf("Max from %v and %v is: %v\n", x, y, maxNum)
}
