package main

/*
84. 柱状图中最大的矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。

图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。

示例:

输入: [2,1,5,6,2,3]
输出: 10
*/
func main() {
	largestRectangleArea2([]int{2, 1, 5, 6, 2, 3})
}

//1.暴力法
func largestRectangleArea(heights []int) int {
	length := len(heights)
	if length == 0 {
		return 0
	}

	max := 0
	for i, v := range heights {
		min := v
		if max < v {
			max = v
		}
		for j := i + 1; j < length; j++ {
			if heights[j] < min {
				min = heights[j]
			}
			temp := (j - i + 1) * min
			if max < temp {
				max = temp
			}
		}
	}
	return max
}

//2.寻找每个柱子的左右边界后计算
func largestRectangleArea2(heights []int) int {
	length := len(heights)
	if length == 0 {
		return 0
	}
	max := 0

	for i, v := range heights {

		left, right := i, i
		for left > 0 {
			if heights[left-1] >= v {
				left--
			} else {
				break
			}
		}

		for right < length-1 {
			if heights[right+1] >= v {
				right++
			} else {
				break
			}
		}
		temp := (right - left + 1) * v
		if max < temp {
			max = temp
		}
	}

	return max
}

//3.单调栈
func largestRectangleArea3(heights []int) int {
	length := len(heights)
	if length == 0 {
		return 0
	}
	//每heights中元素的左右边界
	left, right := make([]int, length), make([]int, length)
	//单调递减s栈
	monoStack := []int{}

	for i := 0; i < length; i++ {
		//清除大于 height[i] 的栈内元素
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
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
	monoStack = []int{}
	for i := length - 1; i >= 0; i-- {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = length
		} else {
			right[i] = monoStack[len(monoStack)-1]
		}
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

//4.单调栈+常数优化
func largestRectangleArea4(heights []int) int {
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
