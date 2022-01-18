package main

import "fmt"

func main() {
    a := []string{"1", "qqq", "ppp"}
    b := []string{"2", "qq12q", "pp12p"}
    c := b
    for _, aa := range  a{
        c = append(c, aa)
    }
    fmt.Println(c)
}
