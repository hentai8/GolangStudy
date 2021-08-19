package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	c := 100
	return c
}

// 匿名多个返回值
func fool2(a string, b int) (int, int) {
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	return 666, 777
}

// 有形參名称的返回多个返回值
func fool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("fool3")
	r1 = 1000
	r2 = 2000
	return
}

func fool4(a string, b int) (r1, r2 int) {
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("fool3")
	r1 = 1000
	r2 = 2000
	return
}

func main() {
	mainc := foo1("abc", 10)
	ret1, ret2 := fool2("miao", 15)
	ret3, ret4 := fool3("miao", 15)
	fmt.Println(mainc)
	fmt.Println(ret1, ret2)
	fmt.Println(ret3, ret4)
}
