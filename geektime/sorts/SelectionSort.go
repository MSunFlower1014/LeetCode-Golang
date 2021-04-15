package main

import "fmt"

/*
Selection Sort 选择排序

排序空间复杂度为 O(1)，是一种原地排序算法。
最好情况时间复杂度、最坏情况和平均情况时间复杂度都为 O(n2)
不稳定算法
*/
func main() {
	list := []int{2, 9, 4, 1, 3}
	SelectionSort(list)
	fmt.Printf("list : %v", list)
}

func SelectionSort(list []int) {
	length := len(list)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		s := i
		v := list[i]
		//选择最小值
		for j := i; j < length; j++ {
			if v > list[j] {
				s = j
				v = list[j]
			}
		}
		//放置到有序区间的末尾
		list[s] = list[i]
		list[i] = v
	}
}
