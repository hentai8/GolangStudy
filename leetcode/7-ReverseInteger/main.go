package main

import "fmt"

func main() {
    fmt.Println(reverserInteger(123))
}

func reverserInteger(x int) int {
    tmp := 0
    for x != 0 {
        tmp = tmp*10 + x%10
        x = x / 10
    }
    // << 表示左移31位
    if tmp > 1<<31-1 || tmp < -(1<<31) {
        return 0
    }
    return tmp
}
