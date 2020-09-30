package main

import "fmt"

/*
75. 颜色分类
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:

输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
*/
func main() {
	list := []int{2, 0, 2, 1, 1, 0, 1, 2}
	sortColors(list)
	fmt.Print(list)
}

func sortColors(nums []int) {
	redIndex := 0
	blueIndex := len(nums) - 1

	for i := range nums {
		if nums[i] == 0 {
			nums[i] = nums[redIndex]
			nums[redIndex] = 0
			redIndex++
			continue
		}
		for nums[i] == 2 && i < blueIndex {
			nums[i] = nums[blueIndex]
			nums[blueIndex] = 2
			blueIndex--
		}
		if nums[i] == 0 {
			nums[i] = nums[redIndex]
			nums[redIndex] = 0
			redIndex++
			continue
			a := make(map[int]int)
		}
	}
}
