package main

import "errors"

func main() {
    var err error
    err = errors.New("this is a demo error")
    println(err.Error())

}
