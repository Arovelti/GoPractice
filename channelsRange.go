package main

import (
	"fmt"
	"time"
)

func fib(n int, c chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		c <- a // blocked until value recieved from channel
		a, b = b, a+b
		time.Sleep(2 * time.Second)
	}
	close(c)
}

func counter(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
	close(c)
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go fib(10, c1)
	go counter(10, c2)

	c1Closed := false
	c2Closed := false

	for {
		select {
		case n, ok := <-c1:
			if !ok {
				c1Closed = true
				if c1Closed && c2Closed {
					return
				}
			} else {
				fmt.Println("Fib: ", n)
			}
		case m, ok := <-c2:
			if !ok {
				c2Closed = true
				if c1Closed && c2Closed {
					return
				}
			} else {
				fmt.Println("Counter: ", m)
			}
		}
	}
}
