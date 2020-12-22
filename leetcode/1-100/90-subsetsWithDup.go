package main

import "sort"

/*
90. 子集 II
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/
func main() {

}

func subsetsWithDup(nums []int) [][]int {
	var result [][]int

	sort.Ints(nums)
	var temp []int
	return add(nums, 0, temp, result)
}

//回溯法
func add(nums []int, start int, temp []int, result [][]int) [][]int {
	newTmp := make([]int, len(temp))
	// 切片底层公用数据，所以要copy
	copy(newTmp, temp)
	result = append(result, newTmp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		result = add(nums, i+1, temp, result)
		temp = temp[:len(temp)-1]
	}
	return result
}
