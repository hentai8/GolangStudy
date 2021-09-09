package main

import (
    "fmt"
    "regexp"
    "strconv"
)

func main() {
    res := "1"
    reg := regexp.MustCompile(`<span>([\d]{7,9}) block`)
    heightstr := reg.FindStringSubmatch(res)
    fmt.Println(heightstr)
    height, err := strconv.ParseInt(heightstr[1], 10, 64)
    if err == nil {
       fmt.Println(height, err)
       return
    } else {
       fmt.Println(err.Error())
    }
}
