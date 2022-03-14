package main

import "fmt"

func main() {
    //s :="{\n  amount: 5000000000n,\n  metadata: [],\n  token_code: {\n    address: '0x00000000000000000000000000000001',\n    module: 'STC',\n    name: 'STC'\n  }\n}\n"
    //buffer := bytes.NewBuffer([]byte(s))

    //s := "0x00f2052a01000000000000000000000000000000000000000000000000000001035354430353544300"
    //fmt.Println([]byte(s))

    num := 5000000000
    s := num/256/256/256
    fmt.Println(s)
}
