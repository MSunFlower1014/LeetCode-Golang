package main

import (
	"fmt"
	"sort"
)

/*
16. 最接近的三数之和
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。
假定每组输入只存在唯一答案。

示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
[0,2,1,-3]
1

提示：

3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4
*/
func main() {
	result := threeSumClosest([]int{0, 2, 1, -3}, 1)
	fmt.Print(result)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	dif := target - result
	if dif < 0 {
		dif = -dif
	}
	for i, x := range nums {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := x + nums[l] + nums[r]
			ans := target - sum
			if ans == 0 {
				return target
			}
			if ans > 0 {
				if dif-ans > 0 {
					dif = ans
					result = sum
				}
				l++
			} else {
				if dif+ans > 0 {
					dif = -ans
					result = sum
				}
				r--
			}
		}
	}
	return result
}
