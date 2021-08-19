package main

import "fmt"

func main() {
	// pair<type:string, value:"hentai8">
	var a string
	a = "hentai8"

	// 万能type
	// pair<type:string,value:"hentai8">
	var allType interface{}
	allType = a

	str, ok := allType.(string)
	fmt.Println(str)
	fmt.Println(ok)
}
