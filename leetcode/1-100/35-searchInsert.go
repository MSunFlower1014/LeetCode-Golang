package main

import "fmt"

/*
35. 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
*/
func main() {
	fmt.Print(searchInsert([]int{1, 3, 5, 6}, 4))
}

func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v < target {
			continue
		}
		return i
	}
	return len(nums)
}
