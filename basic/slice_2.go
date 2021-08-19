package main

import "fmt"

func printArray(myArray []int) {
	// 切片的传递是引用传递,是“指针”
	// _表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value =", value)
	}
	myArray[0] = 100
}

func main() {
	// 动态数组，切片slice，本身就是指针，在传參上非常友好
	myArray := []int{1, 2, 3, 4}
	fmt.Printf("type of myArray = %T\n", myArray)
	printArray(myArray)
	fmt.Println("===========")
	for _, value := range myArray {
		fmt.Println("value =", value)
	}

}
