package main

import (
	"fmt"
	"sort"
)

/*
18. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
注意：
答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

[-3,-1,0,2,4,5]
0
*/
func main() {
	fmt.Print(fourSum([]int{-3, -1, 0, 2, 4, 5}, 0))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		if nums[i] > target/4 {
			break
		}
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if nums[i]+nums[j] > target/2 {
				break
			}
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			m := len(nums) - 1
			for k < m {
				if k != j+1 && nums[k] == nums[k-1] {
					k++
					continue
				}
				if m != len(nums)-1 && nums[m] == nums[m+1] {
					m--
					continue
				}
				if nums[i]+nums[j]+nums[k]+nums[m] > target {
					m--
					continue
				}
				if nums[i]+nums[j]+nums[k]+nums[m] < target {
					k++
					continue
				}
				if k < m && nums[i]+nums[j]+nums[k]+nums[m] == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[m]})
					k++
					m--
				}
			}
		}
	}
	return result
}
