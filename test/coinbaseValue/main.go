package main

import (
    "fmt"
    "math"
    "strconv"
)

func GetCoinbaseValue(height int64) (coinbaseValue string) {
    t := int64(height / 840000)
    fmt.Println(50/(math.Pow(2, float64(t))))
    coinbaseValue = strconv.FormatFloat(50/(math.Pow(2, float64(t))), 'G', -1, 64)
    return
}

func main() {
    h := int64(8400000)
    cv := GetCoinbaseValue(h)
    fmt.Println(cv)
}
