package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	s := "ckb"
	fmt.Print(s)
	s = strings.ToUpper(s)
	fmt.Print(s)
	ss := strings.Split(s,",")
	fmt.Print(ss)
	fmt.Print(reflect.TypeOf(ss))
	//fmt.Print(s)
}
