package main

import (
    "encoding/json"
    "fmt"
    "github.com/asmcos/requests"
)

type Reject struct {
    Time    []string      `json:"time"`
    Rates   []json.Number `json:"rates"`
    Rejects []json.Number `json:"rejects"`
}

type PoolResp []struct {
    Name             string      `json:"name"`
    Hashrate         string      `json:"pool_hashrate"`
    OrphanBlockCount int         `json:"isolated_block_total"`
    Reject           json.Number `json:"rejects"`
}

func main() {
    //var json map[string]interface{}
    getReq := requests.Requests()
    getReq.SetTimeout(20)
    var res PoolResp
    retry := 3
    resp, err := getReq.Get("https://www.dxpool.com/api/pools")
    err = resp.Json(&res)
    if err != nil {
        fmt.Printf("failed 0")
    } else {
        //fmt.Println("------res=", res, "------")
        rejRetry := 3
        for rejRetry > 0 {
            //根据名字获取所有的拒绝率信息
            i := 0
            for _, coin := range res {
                fmt.Println("coin_name: ", coin.Name)
                rejResp, rejErr := getReq.Get("https://www.dxpool.com/api/pools/" + coin.Name + "/hashrates?minutes=60")
                if rejErr != nil {
                    fmt.Printf("failed 1")
                    retry--
                } else {
                    //rejErr = rejResp.Json(&reject)
                    if rejErr != nil {
                        fmt.Printf("failed 2")
                    } else {
                        var reject Reject
                        if err := json.Unmarshal([]byte(rejResp.Text()), &reject); err == nil {
                            //fmt.Println(reflect.TypeOf(reject))
                            //fmt.Println(reject)
                            fmt.Println(reject.Rejects[22])
                            res[i].Reject = reject.Rejects[22]
                            i = i + 1
                        } else {
                            fmt.Println(err)
                        }
                    }
                }
            }
            fmt.Println(res)
        }
    }
}
