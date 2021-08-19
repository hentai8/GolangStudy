package main

import "fmt"

func main() {
	//define a channel
	// channel有同步能力，如果有一个routine还没有接受到数据，会发生一个阻塞，等别的routine中给channel赋值后，数据接收到了再继续执行，如果没人传值，会一直阻塞下去
	// 同时，一个routine在给channel赋值后也会进行阻塞，在另一个routine接收到信息后，才会继续执行，如果没人读取，会一直阻塞下去
	// 意味着一次读写中，赋值和接受的两个routine都会有一次阻塞，一共两个阻塞
	c := make(chan int)

	go func() {
		defer fmt.Println("go routine end!")
		fmt.Println("go routine running...")
		// 将666发送给c
		c <- 666
	}()

	// 从c中接受数据,并赋值给c
	num := <-c
	//num := 1
	fmt.Println("num = ", num)
	fmt.Println("main routine end!")

}
