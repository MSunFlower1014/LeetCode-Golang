package main

import (
	"fmt"
	"sort"
)

/*
56. 合并区间
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: intervals = [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。
*/
func main() {
	nums := [][]int{{1, 3}, {3, 6}}
	fmt.Print(mergeIntervals(nums))
}

func mergeIntervals(intervals [][]int) [][]int {
	//排序
	mIntArray := &IntArray{intervals}
	sort.Sort(mIntArray)
	intervals = mIntArray.mArr

	var result [][]int
	var temp []int
	for _, v := range intervals {
		if len(temp) == 0 {
			temp = v
			continue
		}
		mid := v[0]
		if mid >= temp[0] && mid <= temp[1] {
			if v[1] > temp[1] {
				temp[1] = v[1]
			}
			continue
		}
		result = append(result, temp)
		temp = v
	}
	result = append(result, temp)
	return result
}

type IntArray struct {
	mArr [][]int
}

//IntArray实现sort.Interface接口
func (arr *IntArray) Len() int {
	return len(arr.mArr)
}

func (arr *IntArray) Swap(i, j int) {
	arr.mArr[i], arr.mArr[j] = arr.mArr[j], arr.mArr[i]
}

func (arr *IntArray) Less(i, j int) bool {
	arr1 := arr.mArr[i]
	arr2 := arr.mArr[j]
	return arr1[0] < arr2[0]
}
