package main

import (
    "errors"
    "fmt"
)

func doSomething() error {
    return errors.New("some error occurred")
}


// 不透明错误处理机制
func main() {

    err := doSomething()
    if err != nil {
        fmt.Println(err)
    }

}
