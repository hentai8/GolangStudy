package main

import "fmt"

func main() {
    input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
    output := maxArea(input)
    fmt.Println(output)
}

func maxArea(height []int) int {
    max := 0
    for i := 0; i < len(height)/2; i++ {
        for j := i + 1; j < len(height); j++ {
            area := min(height[i], height[j]) * (j - i)
            if area > max {
                max = area
            }
        }
    }
    return max
}

func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
