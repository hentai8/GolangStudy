package main

import "fmt"

func main() {
    //maxInt := int64(2<<30)
    //fmt.Println(maxInt)
    answer := myAtoiAnswer("    -43")
    fmt.Println(answer)
}

func myAtoiAnswer(s string) int {
    var digits []int
    // int64(2<<30) 意思是2^30,并转为int64格式
    maxInt, signAllowed, whitespaceAllowed, sign := int64(2<<30), true, true, 1
    for _, c := range s {
        if c == ' ' && whitespaceAllowed {
            continue
        }
        if signAllowed {
            if c == '+' {
                signAllowed = false
                whitespaceAllowed = false
                continue
            } else if c == '-' {
                sign = -1
                signAllowed = false
                whitespaceAllowed = false
                continue
            }
        }
        if c < '0' || c > '9' {
            break
        }
        whitespaceAllowed, signAllowed = false, false
        //fmt.Println(int(c-48))
        digits = append(digits, int(c-48))
    }
    var num, place int64
    place, num = 1, 0
    lastLeading0Index := -1
    for i, d := range digits {
        if d == 0 {
            lastLeading0Index = i
        } else {
            break
        }
    }
    if lastLeading0Index < -1 {
        digits = digits[lastLeading0Index+1:]
    }
    var rtnMax int64
    if sign > 0 {
        rtnMax = maxInt - 1
    } else {
        rtnMax = maxInt
    }
    digitsCount := len(digits)
    for i := digitsCount - 1; i >= 0; i-- {
        num += int64(digits[i]) * place
        place *= 10
        if digitsCount-i > 10 || num > rtnMax {
            return int(int64(sign) * rtnMax)
        }
    }
    num *= int64(sign)
    return int(num)
}
