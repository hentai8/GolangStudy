package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// now0 := time.Now()
	//
	// time.Sleep(1000000000)
	//
	// now1 := time.Now()
	//
	// s := int(now1.Sub(now0).Minutes())
	// fmt.Println(s)
	durationString := strconv.Itoa(1*24) + "h"
	day, _ := time.ParseDuration(durationString)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Add(day).Format("2006-01-02 15:04:05"))
	// fmt.Println(FormatRFC3339Time(time.Now()))

}

func FormatRFC3339Time(t string) string {
	t1, _ := time.Parse(time.RFC3339, t)
	t = t1.Format("2006-01-02 15:04:05")
	return t
}
