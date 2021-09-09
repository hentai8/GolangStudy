package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/asmcos/requests"
    "io/ioutil"
    "net/http"
)

type Height struct {
    Result struct {
        Height int64 `json:"count"`
    } `json:"result"`
}

func main() {
    data := requests.Datas{
        "name":    "requests_post_test",
    }
    resp, _ := requests.Post("https://www.httpbin.org/post", data, )
    println(resp.Text())

    //retry := 3
    //for retry > 0 {
    //    getReq := requests.Requests()
    //    getReq.Header.Set("Content-Type", "application/json")
    //    data := requests.Params{
    //        "jsonrpc": "2.0",
    //        "id":      "0",
    //        "method":  "getblockcount",
    //    }
    //    fmt.Println("data:",data)
    //    res, err := getReq.Post("https://wallet.dero.io/json_rpc", data)
    //    if err != nil {
    //        fmt.Printf("failed to get 1 pool data from dxpool account mining, err: %v", err.Error())
    //        retry--
    //    } else {
    //        fmt.Println("res:", res.Text())
    //        var heightStruct Height
    //        if err = json.Unmarshal([]byte(res.Text()), &heightStruct); err == nil {
    //            height := heightStruct.Result.Height
    //            fmt.Println(height)
    //            return
    //        } else {
    //            retry--
    //        }
    //    }
    //}

    //jsonStr, err := json.Marshal("{\"jsonrpc\": \"2.0\",\"id\": \"0\",\"method\": \"getblockcount\"}")
    jsonStr := "{\"jsonrpc\": \"2.0\",\"id\": \"0\",\"method\": \"getblockcount\"}"
    req, err := http.NewRequest("GET", "https://wallet.dero.io/json_rpc", bytes.NewBuffer([]byte(jsonStr)))
    if err != nil {
        return
    }
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return
    }
    defer res.Body.Close()
    response, _ := ioutil.ReadAll(res.Body)
    respp := string(response)
    fmt.Println(respp)
    var heightStruct Height
    if err = json.Unmarshal([]byte(respp), &heightStruct); err == nil {
        h := heightStruct.Result.Height
        fmt.Println(h)
    } else {
        fmt.Println("err:",err)
    }

}
