package main

import "sort"

/*
39. 组合总和
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/
func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(candidates []int, target int) [][]int {
	var ret [][]int
	for i, d := range candidates {
		if target-d < 0 {
			break
		} else if target-d == 0 {
			ret = append(ret, []int{d})
			continue
		}
		for _, v := range dfs(candidates[i:], target-d) {
			ret = append(ret, append([]int{d}, v...))
		}
	}
	return ret
}
