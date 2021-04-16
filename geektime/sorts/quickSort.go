package main

import (
	"fmt"
)

/*
Quick Sort 快速排序

快排的时间复杂度也是 O(nlogn)。
属于原地排序

分区选值优化：
1.从区间的首，尾，中间分别取出一个数，对比大小取中间值作为分区点
2.随机选值做为分区点
*/
func main() {
	list := []int{2, 9, 4, 1, 3}
	QuickSort(list, 0, len(list)-1)
	fmt.Printf("list : %v", list)

}

func QuickSort(list []int, l, r int) {
	if l >= r {
		return
	}
	mid := Partition(list, l, r)
	QuickSort(list, l, mid-1)
	QuickSort(list, mid+1, r)
}

func Partition(list []int, l, r int) int {
	pivot := list[r]
	i := l
	for j := l; j < r; j++ {
		if list[j] < pivot {
			temp := list[j]
			list[j] = list[i]
			list[i] = temp
			i++
		}
	}
	temp := list[r]
	list[r] = list[i]
	list[i] = temp
	return i
}
