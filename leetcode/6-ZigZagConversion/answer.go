package main

import "fmt"

func main() {
    s := "PAYPALISHIRING"
    numRows := 4
    output := convertAnswer(s, numRows)
    fmt.Println(output)
}

func convertAnswer(s string, numRows int) string {
    // make slice 单次长度为numRows,总长度也为numRows
    matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2

    for i := 0; i < len(s); {
        if down != numRows {
            matrix[down] = append(matrix[down], byte(s[i]))
            down++
            i++
        } else if up > 0 {
            matrix[up] = append(matrix[up], byte(s[i]))
            up--
            i++
        } else {
            up = numRows - 2
            down = 0
        }
    }
    solution := make([]byte, 0, len(s))
    for _, row := range matrix {
        for _, item := range row {
            solution = append(solution, item)
        }
    }
    return string(solution)
}
