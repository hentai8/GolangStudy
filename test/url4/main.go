package main

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
    "strconv"
)

func main() {
    rconn, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println("redis.Dial err=", err)
        return
    }
    keys, err := redis.Strings(rconn.Do("keys", "ltc:*"))
    if err != nil {
        fmt.Println("redis.Do err=", err)
        return
    }
    for _, v := range keys {
        height, err := redis.String(rconn.Do("get", v))
        h, err := strconv.ParseInt(height, 10, 64)
        if err != nil {
            fmt.Println("strconv.ParseInt err=", err)
            return
        }
        //fmt.Println("compare:", v, ":", h)
        //fmt.Println("the difference of height", v, ":", h-mainH)
        //// 比较高度,如果高度是1,则tell给targetUrl
        //fmt.Println("!.!.!.start mining empty block")
        h = h
        url := v[4:]
        fmt.Println(url)
        NotifyData, err := redis.String(rconn.Do("get", "ltcNotify:"+url))
        fmt.Println(NotifyData)
    }
}
