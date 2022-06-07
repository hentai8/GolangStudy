package main

import (
    "fmt"
    "strconv"
    "time"
)

func main() {
    t := time.Now()
    fmt.Println(t.UnixNano()/1000)
    fmt.Println(1652950028 - 1650358028)

    var d int64
    d = 100

    updatedAt := time.Now()
    fmt.Println(updatedAt)
    durationString := strconv.FormatInt(d * 24, 10) + "h"
    day, _ := time.ParseDuration(durationString)
    endDate := updatedAt.Add(day)

    fmt.Println(endDate)

}
