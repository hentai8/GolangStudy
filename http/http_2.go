package main


import (
	"encoding/json"
	"fmt"
)

type people struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   int    `json:"id"`
}

type student struct {
	people
	ID int `json:"sid"`
}

func main() {
	msg := "{\"name\":\"zhangsan\", \"age\":18, \"id\":122463, \"sid\":122464}"
	var someOne student
	if err := json.Unmarshal([]byte(msg), &someOne); err == nil {
		fmt.Println(someOne)
		fmt.Println(someOne.people)
	} else {
		fmt.Println(err)
	}
}
