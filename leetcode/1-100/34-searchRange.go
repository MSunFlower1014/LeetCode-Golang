package main

import "fmt"

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/
func main() {
	fmt.Print(searchRange([]int{2, 2}, 2))
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	if n == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	result := -1
	l, r := 0, n-1
	for l >= 0 && l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			result = mid
			break
		}
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if result == -1 {
		return []int{-1, -1}
	}
	left, right := result, result
	for left > 0 && nums[left-1] == target {
		left--
	}
	for right < n-1 && nums[right+1] == target {
		right++
	}
	return []int{left, right}
}
