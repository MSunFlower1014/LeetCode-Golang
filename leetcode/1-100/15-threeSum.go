package main

import (
	"fmt"
	"sort"
)

/*
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/
func main() {
	nums := []int{-1, -1, 1, 0}
	fmt.Print(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			v := nums[i] + nums[l] + nums[r]
			if v < 0 {
				l++
				continue
			} else if v > 0 {
				r--
				continue
			}
			res = append(res, []int{
				nums[i], nums[l], nums[r],
			})
			l++
			r--
			for l < r && nums[l] == nums[l-1] {
				l++
				continue
			}
			for l < r && nums[r] == nums[r+1] {
				r--
				continue
			}
		}
	}
	return res
}
