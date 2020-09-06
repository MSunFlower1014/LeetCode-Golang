package main

/*
57. 插入区间
给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。



示例 1：

输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
示例 2：

输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
*/
import (
	"fmt"
)

func main() {
	nums := [][]int{{11, 12}}
	fmt.Print(insert(nums, []int{5, 9}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	left := newInterval[0]
	right := newInterval[1]
	length := len(intervals)
	if length == 0 {
		result = append(result, newInterval)
		return result
	}
	leftIndex, rightIndex := 0, length-1
	for i := 0; i < length; i++ {
		temp := intervals[i]
		if temp[1] < left {
			leftIndex++
			result = append(result, intervals[i])
		} else {
			break
		}
	}
	for i := length - 1; i >= 0; i-- {
		temp := intervals[i]
		if temp[0] > right {
			rightIndex--
		} else {
			break
		}
	}
	flag := true
	if leftIndex >= length {
		leftIndex = length
		flag = false
	}
	if rightIndex < 0 {
		rightIndex = -1
		flag = false
	}

	if flag {
		if intervals[rightIndex][1] > right {
			right = intervals[rightIndex][1]
		}
		if intervals[leftIndex][0] < left {
			left = intervals[leftIndex][0]
		}
	}

	var mid = []int{left, right}
	result = append(result, mid)

	for i := rightIndex + 1; i < length; i++ {
		result = append(result, intervals[i])
	}
	return result
}
