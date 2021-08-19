package main

import "fmt"

func return0() int {
	fmt.Println("return")
	return 0
}

func defer0() int {
	fmt.Println("defer")
	return 0
}

// defer在return后面
func returnAndDefer() int {
	defer defer0()
	return return0()
}

func main() {
	// defer是在当前函数体结束的时候进行的操作，比如加锁解锁、打开关闭文件
	// defer是堆栈的
	defer fmt.Println("main end 1")
	defer fmt.Println("main end 2")
	fmt.Println("hello go 1")
	fmt.Println("hello go 2")
	returnAndDefer()
}
