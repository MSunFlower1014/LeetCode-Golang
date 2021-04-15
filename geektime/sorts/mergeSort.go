package main

import "fmt"

/*
Merge Sort 归并排序
使用分治思想
将数组分成多个小数组
将多个小数组排序合并后即可

归并排序是一个稳定的排序算法。
归并排序的执行效率与要排序的原始数组的有序程度无关，所以其时间复杂度是非常稳定的，
不管是最好情况、最坏情况，还是平均情况，时间复杂度都是 O(nlogn)。
空间复杂度是 O(n)。
*/
func main() {
	list := []int{2, 9, 4, 1, 3}
	list = MergeSort(list)
	fmt.Printf("list : %v", list)
}

func MergeSort(list []int) []int {
	length := len(list)
	if length <= 1 {
		return list
	}

	mid := length / 2
	left := MergeSort(list[:mid])
	right := MergeSort(list[mid:])
	return Merger(left, right)
}

func Merger(left, right []int) []int {
	lLength := len(left)
	rLength := len(right)
	result := make([]int, lLength+rLength)
	l, r, t := 0, 0, 0
	for l < lLength && r < rLength {
		if left[l] < right[r] {
			result[t] = left[l]
			t++
			l++
		} else {
			result[t] = right[r]
			t++
			r++
		}
	}

	for l < lLength {
		result[t] = left[l]
		t++
		l++
	}
	for r < rLength {
		result[t] = right[r]
		t++
		r++
	}
	return result
}
