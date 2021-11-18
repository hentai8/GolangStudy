package main

import "fmt"

func main() {
    x := 3
    a := 5
    b := 0
    if x > a && (x < b || b == 0) {
        fmt.Println("nmsl")
    }
}
