package main

import "fmt"

func main() {
	//A := big.NewInt(1)
	//B := big.NewInt(2)
	//C := big.NewInt(5)
	//A.Add(A, B)
	//fmt.Println(A)
	//fmt.Println(C)

	s := 1
	ss := fmt.Sprintf("%016d", s)
	fmt.Println(ss)
}
