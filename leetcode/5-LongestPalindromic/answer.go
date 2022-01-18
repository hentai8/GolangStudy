package main

import "fmt"

func main() {
    input := "babad"
    output := longestPalindromeAnswer(input)
    fmt.Println(output)
}

// 马拉车算法
// Manacher Algorithm
// 融合了动态规划和中心扩散
// 终极nb
func longestPalindromeAnswer(s string) (res string) {
    if len(s) < 2 {
        return s
    }
    // rune,用来计算字符长度,会把特殊字符/英文/中文都识别成1
    newS := make([]rune, 0)
    newS = append(newS, '#')
    for _, c := range s {
        newS = append(newS, c)
        newS = append(newS, '#')
    }
    // dp[i]: 以预处理字符串下标 i 为中⼼的回⽂半径(奇数⻓度时不包括中⼼)
    // maxRight: 通过中⼼扩散的⽅式能够扩散的最右边的下标 // center: 与 maxRight 对应的中⼼字符的下标
    // maxLen: 记录最⻓回⽂串的半径
    // begin: 记录最⻓回⽂串在起始串 s 中的起始下标
    dp, maxRight, center, maxLen, begin := make([]int, len(newS)), 0, 0, 1, 0
    for i := 0; i < len(newS); i++ {
        if i < maxRight {
            //关键
            dp[i] = min(maxRight-i,dp[2*center-i])
        }
        //中心扩散法更新
        left, right := i-(1+dp[i]), i+(1+dp[i])
        for left >= 0 && right <len(newS) && newS[left] == newS[right] {
            dp[i]++
            left--
            right++
        }
        // 更新maxRight,它是遍历过的i的i+dp[i]的最大者
        if i + dp[i] > maxRight {
            maxRight = i + dp[i]
            center = i
        }
        // 记录最长会问子串的长度和相应它在原始字符串中的起点
        
    }
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

