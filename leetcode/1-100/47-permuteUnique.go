package main

import (
	"fmt"
	"sort"
)

/*
47. 全排列 II
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/
func main() {
	fmt.Print(permuteUnique([]int{1, 1, 3}))
}

var resultUnique [][]int
var m = make(map[int]int)

func permuteUnique(nums []int) [][]int {
	var pathNums []int
	var used = make([]bool, len(nums))
	sort.Ints(nums)
	// 清空全局数组（leetcode多次执行全局变量不会消失）
	resultUnique = [][]int{}
	m = make(map[int]int)
	backtrackUnique(nums, pathNums, used, true)
	return resultUnique
}

func backtrackUnique(nums, pathNums []int, used []bool, first bool) {
	// 结束条件：走完了，也就是路径上的数字总数等于原始列表总数
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		// 切片底层公用数据，所以要copy
		copy(tmp, pathNums)
		// 把本次结果追加到最终结果上
		m[tmp[0]]++
		resultUnique = append(resultUnique, tmp)
		return
	}

	// 开始遍历原始数组的每个数字
	for i := 0; i < len(nums); i++ {
		if first && m[nums[i]] > 0 {
			continue
		}

		/*
			第三个条件 used[i - 1] 与 !used[i - 1] 都可以完成剪枝
			!used[i - 1] 将只允许重复元素得第一个元素
			used[i - 1] 将只允许重复元素得最后一个元素
		*/
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		// 检查是否访问过
		if !used[i] {
			// 没有访问过就选择它，然后标记成已访问过的
			used[i] = true
			// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
			pathNums = append(pathNums, nums[i])
			backtrackUnique(nums, pathNums, used, false)
			// 撤销刚才的选择，也就是恢复操作
			pathNums = pathNums[:len(pathNums)-1]
			// 标记成未使用
			used[i] = false
		}
	}
}
