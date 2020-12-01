package main

import "fmt"

/**
81. 搜索旋转排序数组 II
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

示例 1:

输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
示例 2:

输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false
*/
func main() {
	var nums []int = []int{2, 5, 6, 0, 0, 1, 2}
	fmt.Println(search81(nums, 0))
}

func search81(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[l] {
			l++
		} else if nums[mid] == nums[r] {
			r--
		} else if nums[mid] < nums[r] {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else if nums[mid] > nums[l] {
			if target < nums[mid] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}
