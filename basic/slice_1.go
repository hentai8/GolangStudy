package main

import "fmt"

func main() {
	// 固定长度的数组
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index =", index)
		fmt.Println("value =", value)
	}

	fmt.Printf("type of myArray1 = %T\n", myArray1)
	fmt.Printf("type of myArray2 = %T\n", myArray2)
}
