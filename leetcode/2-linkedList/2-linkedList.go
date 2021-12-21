package main

import (
    "container/list"
    "fmt"
    "math"
    "reflect"
)

func main() {
    l1 := list.New()
    l1.PushFront(3)
    l1.PushFront(4)
    l1.PushFront(2)
    l2 := list.New()
    l2.PushFront(4)
    l2.PushFront(6)
    l2.PushFront(5)

    addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *list.List, l2 *list.List) {
    l3 := list.New()
    var v1 float64
    var v2 float64
    i := float64(0)
    j := float64(0)
    for element := l1.Back(); element != nil; element = element.Prev() {
        //v1 = v1 + math.Pow(10, i)* float64(element.Value)
        x := element.Value
        //fmt.Println(x)
        fmt.Println(reflect.TypeOf(x))
        t, ok := x.(int)
        //fmt.Println(t)
        if ok {
            q := float64(t)
            v1 = v1 + math.Pow(10, i)*q
        } else {
            println("err")
        }
        i++
    }

    for element := l2.Back(); element != nil; element = element.Prev() {
        //v1 = v1 + math.Pow(10, i)* float64(element.Value)
        x := element.Value
        //fmt.Println(x)
        fmt.Println(reflect.TypeOf(x))
        t, ok := x.(int)
        //fmt.Println(t)
        if ok {
            q := float64(t)
            v2 = v2 + math.Pow(10, j)*q
        } else {
            println("err")
        }
        j++
    }
    fmt.Println(v1)
    fmt.Println(v2)

    v3 := int(v1 + v2)

    var remainder int
    //var digit int
    var quotient int
    remainder = 1
    //digit = 1
    quotient = 1
    for quotient > 0 {
        remainder = v3 % 10
        quotient = v3 / 10
        v3 = quotient
        l3.PushFront(remainder)
    }
    //l3.PushBack(l1.Front().Value+l2.Front().Value)
    for element := l3.Front(); element != nil; element = element.Next() {
        fmt.Println(element.Value)
    }

}

