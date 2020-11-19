package main

import (
	"fmt"
	"sort"
)

/*
80. 删除排序数组中的重复项 II
给定一个增序排列数组 nums ，你需要在 原地 删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。



说明：

为什么返回数值是整数，但输出的答案是数组呢？

请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下：

// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/
func main() {
	var nums = []int{1, 1, 1, 2, 2, 3}
	result := removeDuplicates80(nums)
	fmt.Println(result)
}

func removeDuplicates80(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	result := 0
	dict := make(map[int]int)
	left, right := 0, len(nums)-1

	for left < right {
		if dict[nums[left]] < 2 {
			dict[nums[left]]++
			left++
			result++
		} else {
			temp := nums[right]
			if dict[temp] < 2 {
				nums[right] = nums[left]
				nums[left] = temp

				dict[temp]++
				left++
				right--
				result++
			} else {
				right--
			}
		}
	}
	if dict[nums[left]] < 2 {
		left++
		result++
	}
	slick := nums[0:result]
	sort.Ints(slick)
	return result
}
