package main

import (
    "fmt"
    "time"
)

func main() {
    i := 0
    for i < 10 {
        go func(i int) {
            fmt.Printf("%d", i)
            i++

        }(i)
        i++
    }
    time.Sleep(time.Millisecond)
}
