package main

import (
    "fmt"
)

func main() {
    s := "abcdaefgh"
    //s := "bbbbbb"
    //s := "abcabcbb"
    //a := lengthOfLongestSubstring(s)
    aA1 := lengthOfLongestSubstringAnswer1(s)
    aA2 := lengthOfLongestSubstringAnswer2(s)
    //fmt.Println(a)
    fmt.Println(aA1)
    fmt.Println(aA2)
}

// 垃圾代码
//func lengthOfLongestSubstring(s string) (result int) {
//    m := make(map[string]int)
//    count := 0
//    result = 0
//    for i, ch := range s {
//        sch := string(ch)
//        x, ok := m[sch]
//        if ok && x != -1 {
//            count = i - m[sch]
//            m = make(map[string]int)
//            for k, v := range m {
//                if v < m[sch]{
//                    m[k] = -1
//                }
//            }
//            fmt.Println(m)
//        } else {
//            m[sch] = i
//            if len(m) >= result {
//                result = len(m)
//            }
//            count++
//            fmt.Println(m)
//        }
//    }
//    return
//}

// 位图
// 用一个集合
func lengthOfLongestSubstringAnswer1(s string) int {
    if len(s) == 0 {
        return 0
    }
    // 新建一个集合
    var bitSet [256]bool
    // 宽度 左 右
    result, left, right := 0, 0, 0
    for left < len(s) {
        // 如果集合内存在右边界值，则左边界+1
        if bitSet[s[right]] {
            bitSet[s[left]] = false
            left++
        } else {
            // 如果集合中不存在右边界值，则右边界值+1
            bitSet[s[right]] = true
            right++
        }
        // 宽度 = 左右边界中间
        if result < right-left {
            result = right - left
        }
        // 如果右边界值到达字符串末端 则退出
        if left+result >= len(s) || right >= len(s) {
            break
        }

    }
    return result
}

// 滑动窗口
func lengthOfLongestSubstringAnswer2(s string) int {
    if len(s) == 0 {
        return 0
    }
    var freq [256]int
    result, left, right := 0, 0, -1
    for left < len(s) {
        if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
            freq[s[right+1]-'a']++
            right++
        } else {
            freq[s[left]-'a']--
            left++
        }
        result = max(result, right-left+1)
    }
    return result
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
