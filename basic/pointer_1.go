package main

import "fmt"

func changeValue(p *int) {
	*p = *p * 100
}

func main() {
	a := 1
	changeValue(&a)
	fmt.Println(a)
}
