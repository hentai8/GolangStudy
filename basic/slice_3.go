package main

import "fmt"

func main() {
	//
	slice1 := []int{1, 2, 3}
	slice2 := []int{}

	// 通过make给slice开辟空间，否则会报错
	var slice3 []int
	slice3 = make([]int, 3)

	var slice4 []int

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	//判断一个slice是否为0
	if slice4 == nil {
		fmt.Println("slice是一个空切片")
	} else {
		fmt.Println(slice4)
	}
}
