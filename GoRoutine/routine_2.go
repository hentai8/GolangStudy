package main

import (
	"fmt"
	"time"
)

func main() {
	// 自执行函数
	// 定义后马上执行
	// 用go创建承载一个形參为空，返回值为空的一个函数
	//go func() {
	//	defer fmt.Println("A.defer")
	//
	//	func() {
	//		defer fmt.Println("B.defer")
	//		// 推出当前的goroutine 退出了当前的函数和外面的母函数
	//		// 如果写return的话，只能退出当前函数
	//		runtime.Goexit()
	//		fmt.Println("B")
	//	}()
	//
	//	fmt.Println("A")
	//}()

	go func(a int, b int) bool {
		fmt.Println("a = ", a, "b = ", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
