package main

import (
	"fmt"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Bookp struct {
}

func (this *Bookp) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Bookp) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	// 新建一本书，书拥有ReadBook和WriteBook功能，让读书者使用书，读书者只能读书，让写书者使用书，写书者只能写书
	// b: pair<type:Book, value:book{}地址>
	b := &Bookp{}

	// r: pair<type: , value:>
	var r Reader
	// r:pair<type:Book, value:book{}地址>
	r = b
	r.ReadBook()

	// w: pair<type: , value:>
	var w Writer
	// w:pair<type:Book, value:book{}地址>
	w = b
	w.WriteBook()

}
