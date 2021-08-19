package main

import "fmt"

//声明全局变量，只能用前3种方法，不能:=
var gA int = 121
var gB = 131

//gC := 200

func main() {
	// declare basic
	// 声明一个变量，不赋值，默认的值是0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a = %T\n", a)

	//声明一个变量，赋值
	var b int = 100
	fmt.Println("b=", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Println("bb=", bb)
	fmt.Printf("type of bb = %T\n", bb)

	//声明一个变量，不定义类型
	var c = 1000
	fmt.Println("c=", c)
	fmt.Printf("type of c = %T\n", c)

	// 常用方法 局部变量
	//声明一个变量 :=
	d := 99
	fmt.Println("d=", d)
	fmt.Printf("type of d = %T\n", d)

	f := 99
	fmt.Println("f=", f)
	fmt.Printf("type of f = %T\n", f)

	fmt.Println(gA)
	fmt.Println(gB)

	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println(xx, yy)
	var kk, ll = 100, "abc"
	fmt.Println(kk, ll)

	//
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println(vv, jj)
}
