package main

import "fmt"

func main() {
	//指针！！！都是指针
	s := []int{1, 2, 3}

	s1 := s[0:2]

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)

	//copy函数 将底层数组的slice一起进行拷贝
	s2 := make([]int, 3)
	copy(s2, s[0:2])

	fmt.Println(s2)

	s2[0] = 200

	fmt.Println(s)
	fmt.Println(s2)
}
