package main

import "fmt"

func main(){
	var n int = 30
	fmt.Println("ok1")
	if n > 20 {
		goto label1
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	label1:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}

