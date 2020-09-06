package main

import "fmt"

/*
55. 跳跃游戏
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

示例 1:

输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
示例 2:

输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/
func main() {
	fmt.Print(canJump([]int{2, 0, 0}))
}

func canJump(nums []int) bool {
	length := len(nums) - 1
	if length < 1 {
		return true
	}
	//当前可以达到的最远距离
	max := nums[0]
	for i := 0; i < max; i++ {
		//可达距离超过数组长度
		if nums[i]+i+1 > length {
			return true
		}
		//可达距离超过max，更新max
		if nums[i]+i+1 > max {
			max = nums[i] + i + 1
		}
	}
	return false
}
