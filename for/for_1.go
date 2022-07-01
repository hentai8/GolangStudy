package main

import (
	"fmt"
	"time"
)

func main() {
	duration := time.Second * 1
	timer := time.NewTimer(duration)
	n := 0
	go func() {
		for {
				<-timer.C
				n++
				fmt.Println(n)
				timer.Reset(duration)
			}
	}()
	time.Sleep(60*time.Second)
}
