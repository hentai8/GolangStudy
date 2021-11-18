package main

import (
	"time"
)

func main()  {
	a := make(chan int)

	go func() {
		a <- 1
	}()

	s := <-a
	println(s)
	time.Sleep(10)
}