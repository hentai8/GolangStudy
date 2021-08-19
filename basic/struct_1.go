package main

import "fmt"

// 声明一种新的数据类型 myInt 是int的一个别名
type myInt int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	book.auth = "666"
}

func changeBook2(book *Book) {
	book.auth = "777"
}
func main() {
	var a myInt = 10
	fmt.Println("a =", a)
	fmt.Printf("tpye of a = %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "hentai8"
	fmt.Printf("%v\n", book1)
	changeBook(book1)
	fmt.Printf("%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
