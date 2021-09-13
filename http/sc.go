package main

import (
    "encoding/json"
    "fmt"
    "github.com/asmcos/requests"
    "reflect"
)

func main() {
    type Height []struct {
        LastBlock int64 `json:"lastblock"`
    }

    getReq := requests.Requests()
    getReq.SetTimeout(20)
    res, err := getReq.Get("https://siastats.info/navigator-api/status")
    if err != nil {
        fmt.Printf("failed to get 1 network height from browser, err: %v", err.Error())
    } else {
        var heightStruct Height
        if err = json.Unmarshal([]byte(res.Text()), &heightStruct); err == nil {
            if len(heightStruct) > 0 {
                height := heightStruct[0].LastBlock
                fmt.Println(height)
                fmt.Println(reflect.TypeOf(height))
            }
        } else {
            fmt.Printf("failed to unmarshal 2 network height from browser, err: %v", err.Error())
        }
    }
}
