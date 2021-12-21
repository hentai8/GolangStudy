package main

import (
    "container/list"
    "fmt"
)

func main() {
    fmt.Println("Linked Lists")
    myList := list.New()
    myList.PushBack(3)
    myList.PushFront(5)
    fmt.Println(myList)
    fmt.Println(myList.Front().Value)
    for element := myList.Front(); element != nil; element = element.Next() {
        fmt.Println(element.Value)
    }
}
