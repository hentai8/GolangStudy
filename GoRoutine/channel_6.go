package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i < 10; i++ {
        go func(data int) {
            fmt.Println(data)
        }(i)
    }
    time.Sleep(2 * time.Second)
}
