package main

import (
	"fmt"
)

func Fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			y = x + y
			x = y - x
		case data := <-quit:
			fmt.Println("quit", data)
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	Fibonacci(c, quit)

}
