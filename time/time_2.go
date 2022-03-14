package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now().Second()
    for {
        fmt.Println(time.Now().Second())
        fmt.Println("123")
        time.Sleep(1000000000)
        if time.Now().Second()-t > 3 {
            break
        }
    }
}
