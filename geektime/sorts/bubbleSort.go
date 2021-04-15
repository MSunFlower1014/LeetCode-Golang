package main

import "fmt"

/*
Bubble Sort 冒泡排序
冒泡排序只会操作相邻的两个数据。每次冒泡操作都会对相邻的两个元素进行比较，
看是否满足大小关系要求。如果不满足就让它俩互换。
一次冒泡会让至少一个元素移动到它应该在的位置，重复 n 次，就完成了 n 个数据的排序工作。

空间复杂度为 O(1)，是一个原地排序算法
冒泡排序是稳定的排序算法
时间复杂度就是 O(n2)
*/

func main() {
	list := []int{2, 9, 4, 1, 3}
	BubbleSort(list)
	fmt.Printf("list : %v", list)
}

func BubbleSort(list []int) {
	length := len(list)
	if length <= 1 {
		return
	}

	for i := 0; i < length-1; i++ {
		flag := true
		for j := 0; j < length-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
				flag = false
			}
		}
		//如果无交换发生则已完成排序直接返回
		if flag {
			return
		}
	}
}
