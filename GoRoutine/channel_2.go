package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个带缓冲的channel
	// 有缓存就可以减少阻塞的情况
	// 当channel已经满了，再往里面写数据就会阻塞
	// 当channel为空，从里面取数据也会阻塞
	c := make(chan int, 3)

	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))


	go func() {
		defer fmt.Println("go routine end!")

		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("go routine is running: len(c)=", len(c), "cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("num=", num)
	}

	fmt.Println("main end!")
}
