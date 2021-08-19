package main

import "fmt"

// interface{}是万能的数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	//fmt.Println(arg)

	//interface{} 该如何区分，此时引用的底层数据类型到底是什么
	//给interface{}提供“类型断言”的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string type", value)
	}

}

type Book1 struct {
	auth string
}

func main() {
	book := Book1{"hentai8"}

	myFunc(book)
	myFunc(100)
	myFunc("wuwuwu")

}
