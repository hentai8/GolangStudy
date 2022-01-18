package main

import (
    "fmt"
    "time"
)

func in(ch chan int)  {
    ch <- 1
}

func out(ch chan int)  {
    // 遍历接收通道数据
    for data := range ch {
        // 打印通道数据
        fmt.Println(data)
        // 当遇到数据0时, 退出接收循环
        if data == 0 {
            break
        }
    }
}

func main() {
    // 构建一个通道
    ch := make(chan int)

    go in(ch)
    go out(ch)

    time.Sleep(5*time.Second)
    //// 开启一个并发匿名函数
    //go func() {
    //    // 从3循环到0
    //    for i := 3; i >= 0; i-- {
    //        // 发送3到0之间的数值
    //        ch <- i
    //        // 每次发送完时等待
    //        time.Sleep(time.Second)
    //    }
    //}()
    //// 遍历接收通道数据
    //for data := range ch {
    //    // 打印通道数据
    //    fmt.Println(data)
    //    // 当遇到数据0时, 退出接收循环
    //    if data == 0 {
    //        break
    //    }
    //}
}