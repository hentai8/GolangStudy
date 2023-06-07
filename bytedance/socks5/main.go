package main

import (
	"bufio"
	"log"
	"net"
)

// nc 127.0.0.1 1080

func main() {
	server, err := net.Listen("tcp", ":1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err.Error())
			continue
		}
		go process(client)
	}
}

// 一个发送啥回复啥的服务
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	ver, err := reader.ReadByte()
}
