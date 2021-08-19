package main

import "fmt"

// 可以在const()添加一个关键字iota，每行的iota都会累加1,第一行的iota默认值是0
// iota只能配合const()一起使用,定义枚举类型，code码
const (
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, j
)

func main() {
	// 通过const定义枚举类型
	const length int = 10

	fmt.Println(length)
	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
	fmt.Println(SHENZHEN)

	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
	fmt.Println(g, h)
	fmt.Println(i, j)
}
