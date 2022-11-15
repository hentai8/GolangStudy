package main

import "fmt"

func main() {
	randomness := 0x01
	i := 0
	for {
		i++
		randomness = randomness + 0x01
		if i > 100 {
			break
		}
	}
	randomnessA, _ := fmt.Printf("%v", string(randomness))
	fmt.Println(randomnessA)
	//fmt.Println(string(randomnessA))
}
