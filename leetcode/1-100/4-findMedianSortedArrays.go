package main

import "fmt"

/*
4. 寻找两个正序数组的中位数
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

数组长度和如为奇数，取1/2 +1 ，如为偶数，取（1/2+（1/2+1））/2
*/
func main() {
	result := findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5})
	fmt.Printf("%f", result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	size := 1
	result := 0
	if length%2 == 0 {
		size = 2
	}
	deleteSum := size % 2
	for i := 1; i < (length/2)+deleteSum; i++ {
		if len(nums1) == 0 {
			nums2 = nums2[1:]
			continue
		}
		if len(nums2) == 0 {
			nums1 = nums1[1:]
			continue
		}
		if nums1[0] < nums2[0] {
			nums1 = nums1[1:]
			continue
		} else {
			nums2 = nums2[1:]
			continue
		}
	}

	for j := 0; j < size; j++ {
		if len(nums1) == 0 {
			result = result + nums2[0]
			nums2 = nums2[1:]
			continue
		}
		if len(nums2) == 0 {
			result = result + nums1[0]
			nums1 = nums1[1:]
			continue
		}
		if nums1[0] < nums2[0] {
			result = result + nums1[0]
			nums1 = nums1[1:]
			continue
		} else {
			result = result + nums2[0]
			nums2 = nums2[1:]
			continue
		}
	}
	return float64(result) / float64(size)
}
