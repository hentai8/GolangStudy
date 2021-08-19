package main

import "fmt"

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("````````````")
		fmt.Println("key=", key)
		fmt.Println("value=", value)
		fmt.Println("````````````")
	}
}

func main() {
	cityMap := make(map[string]string)

	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "Washington"

	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}

	fmt.Println("---------------------")

	delete(cityMap, "Japan")

	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}

	fmt.Println("---------------------")

	cityMap["USA"] = "Trump"

	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}

	printMap(cityMap)
}
