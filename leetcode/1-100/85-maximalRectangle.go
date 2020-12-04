package main

import "fmt"

/*
85. 最大矩形
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例 1：

输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。
*/
func main() {
	matrix := make([][]byte, 4)
	matrix[0] = []byte{'1', '0', '1', '0', '0'}
	matrix[1] = []byte{'1', '0', '1', '1', '1'}
	matrix[2] = []byte{'1', '1', '1', '1', '1'}
	matrix[3] = []byte{'1', '0', '0', '1', '0'}

	/*
		示例的dp为 -- 即为每列的最高有效矩形高度，转为84问题，单调栈解决
		[
		[1 1 1 1]
		[0 0 2 0]
		[1 1 3 0]
		[0 2 4 1]
		[0 3 5 0]
		]

	*/
	maximalRectangle(matrix)
}

func maximalRectangle(matrix [][]byte) int {
	height := len(matrix)
	if height == 0 {
		return 0
	}
	width := len(matrix[0])
	dp := make([][]int, width)
	//统计有效矩形的高度
	for i := 0; i < width; i++ {
		dp[i] = make([]int, height)
		for j := 0; j < height; j++ {
			if matrix[j][i] == '1' {
				if i == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j] + 1
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	fmt.Println(dp)
	max := 0
	for _, v := range dp {
		temp := largestRectangleArea85(v)
		if temp > max {
			max = temp
		}
	}
	return max
}

func largestRectangleArea85(heights []int) int {
	length := len(heights)
	if length == 0 {
		return 0
	}
	//每heights中元素的左右边界
	left, right := make([]int, length), make([]int, length)
	//单调递减s栈
	monoStack := []int{}
	for i := 0; i < length; i++ {
		right[i] = length
	}
	for i := 0; i < length; i++ {
		//清除大于 height[i] 的栈内元素
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			right[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}
		//如果栈为空，则之前元素都大于 height[i] ，左边界为-1
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			//首个小于 height[i] 的值的索引为 monoStack[len(monoStack)-1] ，即为左边界
			left[i] = monoStack[len(monoStack)-1]
		}
		//入栈
		monoStack = append(monoStack, i)
	}

	max := 0
	for i := 0; i < length; i++ {
		temp := (right[i] - left[i] - 1) * heights[i]
		if max < temp {
			max = temp
		}
	}
	return max
}
