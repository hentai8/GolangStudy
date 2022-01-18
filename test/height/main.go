package main

import (
    "encoding/hex"
    "fmt"
    "math"
)

func main() {
    heightBuff := "03d71f21"
    heightLength, err := hex.DecodeString(heightBuff[0:2])
    if err != nil {
        fmt.Println("decode height error=", err.Error())
    }
    fmt.Println(heightLength[0])
    heightLengthInt := int(heightLength[0])
    height := 0
    for i := 0; i < heightLengthInt; i++ {
        heightPlus, err := hex.DecodeString(heightBuff[2+2*i : 4+2*i])
        if err != nil {
            fmt.Println("decode height error=", err.Error())
        }
        height = height + int(heightPlus[0])*int(math.Pow(256, float64(i)))
    }
    fmt.Println("height:", height)


    heightBuff1 := "03d71f21"
    height1, _ := hex.DecodeString(heightBuff1[2:4])
    height2, _ := hex.DecodeString(heightBuff1[4:6])
    height3, _ := hex.DecodeString(heightBuff1[6:8])

    h1 := int64(height1[0])
    h2 := int64(height2[0])
    h3 := int64(height3[0])
    h := h1 + h2*256 + h3*256*256
    fmt.Println("h:", h)

}
