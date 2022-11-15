package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	next := now.Add(time.Second)
	next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
	timer := time.NewTimer(next.Sub(now))
	x := 0
	for {
		select {
		case <-timer.C:
			fmt.Println(x)
			x++
		}
		now := time.Now()
		next := now.Add(time.Second)
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
		timer.Reset(next.Sub(now))
	}
}
