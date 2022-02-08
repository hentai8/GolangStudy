package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    netListen, err := net.Listen("tcp", "127.0.0.1:1024")
    if err != nil {
        fmt.Println("create server err=", err)
        return
    }
    defer netListen.Close()
    fmt.Println("Waiting for clients ...")

    //RedisClient := client

    //defer rconn.Close()
    //Notify, err := redis.String(a.RedisClient.Do("get", "ltcNotify:"+url))
    //Notify, err := redis.String(a.RedisClient.Do("get", "ltcNotify:ltc.ss.dxpool.net:3333"))

    //Notify, err := redis.String(a.RedisClient.Do("get", "ltcNotify:127.0.0.1:3032"))

    //fmt.Println(Notify)

    //等待客户端访问
    for {
        //监听接收
        conn, err := netListen.Accept()
        //如果发生错误，继续下一个循环。
        if err != nil {
            continue
        }
        fmt.Println("tcp connect success")

        // 汇报挖空块
        go HandleConnection(conn)

    }
}

func HandleConnection(conn net.Conn) {
    buffer := make([]byte, 2048) //建立一个slice
    for {
        //读取客户端传来的内容
        n, err := conn.Read(buffer)
        // 当远程客户端连接发生错误（断开）后，终止此协程。
        if err != nil {
            fmt.Println("connection error: ", err)
            return
        }
        fmt.Println("receive data string:\n", string(buffer[:n]))

        if string(buffer[:n]) != "startPoolWatcher" {
            continue
        }

        for {
            fmt.Println("OK!")
            _, err := conn.Write([]byte("OK!"))
            if err != nil {
                fmt.Println("tell err:", err.Error())
                break
            }
            time.Sleep(5 * time.Second)
        }

        ////返回给客户端的信息
        //for data := range ltcData {
        //    fmt.Println("SEND:", data)
        //    str := fmt.Sprintf("%v", data)
        //    _, err := conn.Write([]byte(str))
        //    if err != nil {
        //        fmt.Println("tell err:", err.Error())
        //        break
        //    }
        //}
    }
}
