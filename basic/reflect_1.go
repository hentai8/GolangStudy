package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("type:", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 3.141592654
	reflectNum(num)
}
