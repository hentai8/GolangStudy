package main

import "fmt"

func main() {
    n := []int{2, 7, 11, 15}
    t := 9
    a := twoSum(n, t)
    aA := twoSum(n, t)
    fmt.Println(a)
    fmt.Println(aA)
}

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        answer := target - nums[i]
        for j := 0;j< len(nums); j++{
            if answer == nums[j]{
                return []int{i, j}
            }
        }
    }
    return nil
}

func twoSumAnswer(nums []int, target int) []int {
    // 定义一个map集合 包含key和value
    m := make(map[int]int)
    for i := 0; i< len(nums);i++{
        another := target - nums[i]
        //在m中寻找是否有m[another]，如果有，就输出，如果没有，就设置m[another]为i
        if _, ok := m[another]; ok{
            return []int{m[another], i}
        }
        m[nums[i]] = i
    }
    return nil
}
