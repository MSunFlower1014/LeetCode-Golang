package main

import "fmt"

/*
Binary Search 二分法查找
时间复杂度 -  O(logn)

二分查找依赖的是顺序表结构，简单点说就是数组
二分查找针对的是有序数据

如果数据之间的比较操作非常耗时，不管数据量大小，我都推荐使用二分查找。
需要尽可能地减少比较次数，而比较次数的减少会大大提高性能，这个时候二分查找就比顺序遍历更有优势。

二分查找不需要其他额外的空间
*/
func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := BinarySearch(list, 3)
	fmt.Printf(" result is %v", result)
}

func BinarySearch(list []int, n int) int {
	r := len(list) - 1
	if r == 0 || n > list[r] {
		return -1
	}
	l := 0
	for l <= r {
		mid := l + (r-l)>>1
		temp := list[mid]
		if temp == n {
			return mid
		}
		if temp < n {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
