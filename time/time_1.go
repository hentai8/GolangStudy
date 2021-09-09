package main

import (
    "fmt"
    "time"
)

func main() {
    // timer是一个channel
    timer := time.NewTimer(3 * time.Second)
    fmt.Println("当前时间为:", time.Now())
    go func() {
        for {
            //fmt.Println("当前时间为:", time.Now())
            t := <-timer.C
            fmt.Println("当前时间为:", t)
            timer.Reset(3 * time.Second)
        }
    }()
    //重置定时器时间

    for {
        time.Sleep(1 * time.Second)
    }
}
