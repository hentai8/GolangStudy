package main

import (
    "errors"
    "fmt"
)

// 哨兵模式
// is
// erros.Is()会沿着该包装错误所在错误链(error chain)与链上所有被包装的错误进行比较，直到找到一个匹配的错误
var ErrSentinel = errors.New("the underlying sentinel error")

func main() {
    err1 := fmt.Errorf("wrap err1: %w", ErrSentinel)
    err2 := fmt.Errorf("wrap err1: %w", err1)
    if errors.Is(err2, ErrSentinel) {
        println("err is ErrSentinel")
        return
    }
    println("err is not ErrSentinel")

}
