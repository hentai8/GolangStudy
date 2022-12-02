package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	// 尝试与 google.com:80 建立 tcp 连接
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close() // 退出关闭连接

	// 通过连接发送 GET 请求，访问首页
	_, err = fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	if err != nil {
		log.Fatal(err)
	}

	dat, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dat))
}
