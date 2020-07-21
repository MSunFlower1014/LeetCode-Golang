package main

import (
	"fmt"
	"sort"
)

/*
41. 缺失的第一个正数
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。



示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1
*/
func main() {
	fmt.Print(firstMissingPositive2([]int{1, 2, 0}))
}

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	target := 1
	for _, v := range nums {
		if v <= 0 {
			continue
		}
		if v < target {
			continue
		}
		if v == target {
			target++
			continue
		}
		break
	}
	return target
}

func firstMissingPositive2(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
