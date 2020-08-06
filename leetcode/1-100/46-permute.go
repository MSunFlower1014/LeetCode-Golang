package main

import "fmt"

/*
46. 全排列
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

result = []
func backtrack(路径，选择列表) {
	if 满足结束条件 {
		result.add(路径)
	}
	return

	for 选择 in 选择列表 {
		做选择
		backtrack(路径，选择列表)
		撤销选择
	}
}

作者：DCCooper
链接：https://leetcode-cn.com/problems/permutations/solution/gojie-fa-xiang-xi-zhu-shi-hui-su-mo-ban-zou-tian-x/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func main() {
	fmt.Print(permute([]int{1, 2, 3}))
}

// 最终结果
var result [][]int

// 回溯核心
// nums: 原始列表
// pathNums: 路径上的数字
// used: 是否访问过
func backtrack(nums, pathNums []int, used []bool) {
	// 结束条件：走完了，也就是路径上的数字总数等于原始列表总数
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		// 切片底层公用数据，所以要copy
		copy(tmp, pathNums)
		// 把本次结果追加到最终结果上
		result = append(result, tmp)
		return
	}

	// 开始遍历原始数组的每个数字
	for i := 0; i < len(nums); i++ {
		// 检查是否访问过
		if !used[i] {
			// 没有访问过就选择它，然后标记成已访问过的
			used[i] = true
			// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
			pathNums = append(pathNums, nums[i])
			backtrack(nums, pathNums, used)
			// 撤销刚才的选择，也就是恢复操作
			pathNums = pathNums[:len(pathNums)-1]
			// 标记成未使用
			used[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	var pathNums []int
	var used = make([]bool, len(nums))
	// 清空全局数组（leetcode多次执行全局变量不会消失）
	result = [][]int{}
	backtrack(nums, pathNums, used)
	return result
}
