package main

import "fmt"

/*
45. 跳跃游戏 II
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例:

输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
*/
func main() {
	fmt.Print(jump2([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	length := len(nums)
	position := nums[length-1]
	step := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
				step++
				break
			}
		}
	}
	return step
}

func jump2(nums []int) int {
	length := len(nums)
	step := 0
	maxPosition := 0
	end := 0

	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			step++
			end = maxPosition
		}
	}
	return step
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
