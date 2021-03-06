package main

import (
	"fmt"
	"strconv"
)

/*
60. 第k个排列
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
*/
func main() {
	fmt.Print(getPermutation(4, 4))
}

func getPermutation(n int, k int) string {
	var factorials = []int{1}
	var nums = []string{"1"}
	factorials[0] = 1
	for i := 1; i < n; i++ {
		factorials = append(factorials, factorials[i-1]*i)
		nums = append(nums, strconv.Itoa(i+1))
	}

	k--

	var output []string
	for i := n - 1; i > -1; i-- {
		idx := k / factorials[i]
		k -= idx * factorials[i]
		output = append(output, nums[idx])
		nums = append(nums[0:idx], nums[idx+1:]...)
	}

	result := ""
	for _, v := range output {
		result = result + v
	}
	return result
}
