package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主 goroutine
func main() {
	// 创建一个go程 去执行newTask() 流程
	//go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

//// 主 goroutine
//func main() {
//	// 创建一个go程 去执行newTask() 流程
//	go newTask()
//
//	fmt.Println("main Goroutine exit")
//}
