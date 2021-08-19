package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			c <- i
		}

		// close 可以关闭一个channel
		close(c)
	}()

	//for {
	//	// ok如果为true表示channel没有关闭
	//	if data, ok := <- c; ok{
	//		fmt.Println(data)
	//	} else {
	//		break
	//	}
	//}

	// 与上面相同,可以使用range来迭代不断操作channel
	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("MAIN END")
}
