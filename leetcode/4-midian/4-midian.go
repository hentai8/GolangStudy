package main

import "fmt"

func main() {
    //nums1 := []int{1, 2, 3, 9}
    //nums2 := []int{4, 6, 7, 8}

    nums1 := []int{5, 6, 7, 8}
    nums2 := []int{1, 2, 3, 4}
    //a := findMedianSortedArrays(nums1, nums2)
    a := findMedianSortedArraysAnswer(nums1, nums2)
    fmt.Println(a)

    x := 10
    y := x << 1
    z := 19 & 1
    fmt.Println(y)
    fmt.Println(z)


}

// 先要判断字符串的长短

func findMedianSortedArrays(nums1 []int, nums2 []int) (result float64) {
    l1 := len(nums1)
    l2 := len(nums2)
    var l3 []int
    lAll := int((l1 + l2) / 2)
    fmt.Println(lAll)
    j := 0
    i := 0
    for i < l1 {
        //fmt.Println("!")
        //fmt.Println(l3)
        if j > l2-1 {
            l3 = append(l3, nums1[i])
            i++
        } else if i > l1-1 {
            l3 = append(l3, nums2[j])
            j++
        } else if nums2[j] <= nums1[i] {
            l3 = append(l3, nums2[j])
            j++

        } else {
            l3 = append(l3, nums1[i])
            i++
        }
        fmt.Println(l3)
        fmt.Println(l3[i+j-1])
        fmt.Println(l3[i+j-2])
        if i+j > lAll {
            if i+j-1 == lAll {
                if i > l1-1 {
                    result = float64((l3[i+j-1] + l3[i+j-2]) / 2)
                    fmt.Println("!")
                    return
                }
                if j > l2-1 {
                    result = float64((l3[i+j-1] + l3[i+j-2]) / 2)
                    fmt.Println("!")
                    return
                }
                if nums2[j] <= nums1[i] {
                    result = float64((l3[i+j-1] + l3[i+j-2]) / 2)
                    fmt.Println("!")
                    return
                } else {
                    result = float64((l3[i+j-1] + l3[i+j-2]) / 2)
                    fmt.Println("!!")
                    return
                }
            } else {
                result = float64(l3[i+j-1])
                fmt.Println("!!!")
                return
            }
        }
    }
    fmt.Println("?")
    return
}

func findMedianSortedArraysAnswer(nums1 []int, nums2 []int) (result float64) {
    //假设nums1 的长度小，如果不是，则颠倒一下位置
    if len(nums1) > len(nums2) {
        return findMedianSortedArraysAnswer(nums2, nums1)
    }
    low := 0
    // high为nums1的长度
    high := len(nums1)
    // 2进制向右位移1位，即乘2
    k := (len(nums1) + len(nums2) + 1) >> 1
    nums1Mid := 0
    nums2Mid := 0
    for low <= high {
        // nums1Mid 为 nums1
        nums1Mid = low + (high-low)>>1
        nums2Mid = k - nums1Mid
        if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid-1] {
            high = nums1Mid - 1
        } else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
            low = nums1Mid + 1
        } else {
            break
        }
    }
    midLeft, midRight := 0, 0
    if nums1Mid == 0 {
        midLeft = nums2[nums2Mid-1]
    } else if nums2Mid == 0 {
        midLeft = nums1[nums1Mid-1]
    } else {
        midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
    }
    // 判断(len(nums1)+len(nums2))的奇偶
    // 和1作与，二进制最后一位进行判断，1&1=1，0&1=0，
    if (len(nums1)+len(nums2))&1 == 1 {
        return float64(midLeft)
    }
    if nums1Mid == len(nums1) {
        midRight = nums2[nums2Mid]
    } else if nums2Mid == len(nums2) {
        midRight = nums1[nums1Mid]
    } else {
        midRight = min(nums1[nums1Mid], nums2[nums2Mid])
    }
    return float64(midLeft+midRight) / 2
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

