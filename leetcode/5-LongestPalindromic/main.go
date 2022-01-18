package main

import "fmt"

func main() {
    input := "babad"
    output := longestPalindrome(input)
    fmt.Println(output)
}

// 中心扩散法
func longestPalindrome(s string) (res string) {
    res = ""
    for i := 0; i < len(s); i++ {
        // 奇数判断，即对称轴为字符，(i+i)/2
        res = maxPalindrome(s, i, i, res)
        // 偶数判断，即对称轴为字符之间的间隔，(i+i+1)/2
        res = maxPalindrome(s, i, i+1, res)
    }
    return res
}

// 输入包括完整的字符串，i和j为初始的左右边界，决定对称轴，对称轴为(i+j)/2，res为已得到的最大字符串
func maxPalindrome(s string, i int, j int, res string) (out string) {
    // 初始化字符串
    sub := ""
    // 每次判断字符串的最左和最右是否相同
    for i >= 0 && j < len(s) && s[i] == s[j] {
        // 如果相同 覆盖原来的答案
        sub = s[i : j+1]
        // 向外扩散一格
        i--
        j++
    }
    // 如果输出的字符串长度比输入的要长，则输出新的字符串
    if len(res) < len(sub) {
        out = sub
        return
    }
    // 否则输出原先的字符串
    out = res
    return
}
