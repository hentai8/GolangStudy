package main

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)

func main() {
    c, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println("Connect to redis error", err)
        return
    }
    defer c.Close()

    _, err = c.Do("LPUSH", "list", "one")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }

    username, err := redis.StringMap(c.Do("LRANGE", "list", "0", "-1"))
    if err != nil {
        fmt.Println("redis get failed:", err)
    } else {
        fmt.Printf("Get mykey: %v \n", username)
    }
}
