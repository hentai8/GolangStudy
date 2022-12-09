package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "54.84.144.248:9090")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)

	go func() {
		for {
			dat, _, err := reader.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("sever:", string(dat))
		}
	}()
	for {
		//_, err := conn.Write([]byte("hello\n"))
		time.Sleep(1 * time.Second)

		subscribe := "{\"jsonrpc\":\"2.0\",\"id\": 1, \"method\": \"mining.subscribe\", \"params\": [\"gsa0/1.0\", \"AleoStratum/2.0.0\", null]}\n"

		_, err := conn.Write([]byte(subscribe))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("send subscribe")

		//dat, _, err := reader.ReadLine()
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Println("sever:", string(dat))

		time.Sleep(1 * time.Second)

		authorize := "{\"jsonrpc\":\"2.0\",\"id\": 1, \"method\": \"mining.authorize\", \"params\": [\"huang93.hmc\", \"123\"]}\n"

		_, err = conn.Write([]byte(authorize))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("send authorize")

		time.Sleep(5 * time.Second)

		submit := "{\"jsonrpc\":\"2.0\",\"id\": 1, \"method\": \"mining.submit\", \"params\": [\"aleo1p4jjees0xj2m6hxlgej776elcm74lwlj4m3xuqk9f6zq64zjwgpslkwd42\", \"2298959d\", \"2298959d\", \"COMMITMENT\", \"PROOF\"]}\n"

		_, err = conn.Write([]byte(submit))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("send submit")

		time.Sleep(10 * time.Second)
	}
}
