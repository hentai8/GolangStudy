package main

import (
    "encoding/json"
    "fmt"
    "github.com/asmcos/requests"
    "strconv"
)

func main() {
    type Height struct {
        Data []struct {
            Attributes struct {
                Number string `json:"number"`
            } `json:"attributes"`
        } `json:"data"`
    }

    //type Height struct {
    //    Data []struct {
    //        Height int64 `json:"id"`
    //    } `json:"data"`
    //}

    //type Height map[string]struct {
    //    Id int64 `json:"id"`
    //}
    ////`json:"data"`

    getReq := requests.Requests()
    getReq.SetTimeout(20)
    getReq.Header.Set("Content-Type", "application/vnd.api+json")
    getReq.Header.Set("Accept", "application/vnd.api+json")
    res, err := getReq.Get("https://api.explorer.nervos.org/api/v1/blocks")
    if err != nil {
        fmt.Printf("failed to get 1 network height from browser, err: %v", err.Error())
    } else {
        var heightStruct Height
        fmt.Println(res.Text())
        if err = json.Unmarshal([]byte(res.Text()), &heightStruct); err == nil {
            //if len(heightStruct.Data) > 0 {
            heightsrt := heightStruct.Data[0].Attributes.Number
            height, err := strconv.ParseInt(heightsrt, 10, 64)
            //    height := heightStruct.Data[0].Height
            fmt.Println(height, err)
            //}
        } else {
            fmt.Printf("failed to unmarshal 2 network height from browser, err: %v", err.Error())
        }
    }
}
