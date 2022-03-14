package main

import (
    "fmt"
    "github.com/robertkrimen/otto"
    "io/ioutil"
)

func main() {
    //    vm := otto.New()
    //    vm.Run(`
    //    a = 2+2;
    //    console.log(a);
    //`)
    //    value, err := vm.Get(`
    //    xxx.length
    //`)
    //    if err != nil {
    //        fmt.Println("err")
    //        return
    //    }
    //    fmt.Println(value)

    filePath := "./test.js"

    bytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        panic(err)
    }

    vm := otto.New()

    _, err = vm.Run(string(bytes))
    if err != nil{
        panic(err)
    }

    data := "0x00f2052a01000000000000000000000000000000000000000000000000000001035354430353544300"
    value, err := vm.Call("decodeDepositEventEventData", nil, data)
    if err != nil {
        panic(err)
    }
    fmt.Println(value.String())
}
