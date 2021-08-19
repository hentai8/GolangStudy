package main

import "fmt"

func main() {
	// 声明myMap1是一种map类型，key是string，value是string
	// map和slice一样是指针
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1是一个空map")
	}

	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"
	fmt.Println(myMap1)

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"one":   "sumsung",
		"two":   "xiaomi",
		"three": "iphone",
	}
	fmt.Println(myMap3)
}
