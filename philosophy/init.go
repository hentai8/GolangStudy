package main

import (
    "fmt"
)

func main() {
    m := map[int]string{1: "hello", 2: "hentai8", 3: "miao"}
    n := map[string]string{"h": "hello", "h8": "hentai8", "m": "miao"}
    f := [...]float64{-1,1:0.5}

    fmt.Println(m)
    fmt.Println(n)
    fmt.Println(f)


}
