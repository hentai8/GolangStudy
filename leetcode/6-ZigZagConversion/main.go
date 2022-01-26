package main

import "fmt"

func main() {
    s := "PAYPALISHIRING"
    numRows := 4
    output := convert(s, numRows)
    fmt.Println(output)
}

func convert(s string, numRows int) string {
    l := len(s)
    fmt.Println("length:", l)
    cycle := numRows*2 - 2
    completeCycle := int(l / cycle)
    fmt.Println(completeCycle)
    width := 0
    fmt.Println(completeCycle * cycle)
    if l-completeCycle*cycle <= numRows {
        width = completeCycle*(numRows-1) + 1
        fmt.Println(1)
    } else if l-completeCycle*cycle > numRows {
        width = completeCycle*(numRows-1) + l - completeCycle*cycle - numRows + 1
        fmt.Println(2)
    }
    fmt.Println("width:", width)

    matrix := make([][]string, numRows, numRows)
    Divisor := numRows*2 - 2

    for i := 0; i < width; i++ {
        for j := 0; j < numRows; j++ {
            //if Divisor*i+j < l {
            if j == 0 || j == numRows-1 {
                if Divisor*i+j < l {
                    matrix[j] = append(matrix[j], string(s[Divisor*i+j]))
                }
            } else {
                if Divisor*i-j > 0 && Divisor*i-j < l {
                    matrix[j] = append(matrix[j], string(s[Divisor*i-j]))
                }
                if Divisor*i+j < l {
                    matrix[j] = append(matrix[j], string(s[Divisor*i+j]))
                }
            }
            //}
        }
    }

    answer := ""
    fmt.Println(matrix)

    for _, line := range matrix {
        for _, v := range line {
            answer = answer + v
        }
    }

    return answer
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
