package main

import (
	"fmt"
	"gopkg.in/redis.v5"
    "strconv"
)
func Del(client *redis.Client, key string) *redis.IntCmd {
    intCmd := client.Del(key)
    return intCmd
}

func Get(client *redis.Client, key string) *redis.StringCmd {
    cmd := redis.NewStringCmd("GET", key)
    client.Process(cmd)
    return cmd
}

func Set(client *redis.Client, key string, value string) *redis.StringCmd {
    cmd := redis.NewStringCmd("SET", key, value)
    client.Process(cmd)
    return cmd
}

func Keys(client *redis.Client, key string) *redis.StringSliceCmd {
    stringSliceCmd := client.Keys(key)
    return stringSliceCmd
}

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    x := 1111
    s1 := strconv.FormatInt(int64(x), 10)
    _, err := Set(client, "k1", s1).Result()
    if err != nil {
       fmt.Println("redis.Do err=", err)
       return
    }



    ////client.Keys("ltc:*")
    //keys, err := Keys(client, "ltc:*").Result()
    //if err != nil {
    //    fmt.Println("redis.Do err=", err)
    //    return
    //}
    //fmt.Println(keys)
    //for _, key := range keys {
    //    fmt.Println(key)
    //    _, err := Del(client, key).Result()
    //    if err != nil {
    //        fmt.Println("redis del pre ltc data error=", err.Error())
    //        return
    //    }
    //}
}
