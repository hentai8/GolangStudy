package main

import (
    "errors"
    "fmt"
    "github.com/asmcos/requests"
    "regexp"
    "strconv"
)

func main() {
    retry := 3
    for retry > 0 {
        req := requests.Requests()
        req.SetTimeout(5)
        res, err := req.Get("https://explorer.evolution-network.org/")
        if err != nil {
            fmt.Printf("failed to get 1 network height from browser, err: %v", err.Error())
            retry--
        } else {
            fmt.Println(res)
            reg := regexp.MustCompile(`of ([\d]{7,9}) block`)
            heightstr := reg.FindAllString(res.Text(), -1)
            height, err := strconv.ParseInt(heightstr[0], 10, 64)
            if err == nil {
                fmt.Println(height, err)
                return
            } else {
                retry --
            }

        }
    }
    err := errors.New("failed to get net block height after 3 times")
    fmt.Println("err:", err)
}
