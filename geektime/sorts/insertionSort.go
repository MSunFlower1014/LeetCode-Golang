package main

import "fmt"

/*
InsertionSort 插入排序

空间复杂度是 O(1)，是一个原地排序算法
插入排序是稳定的排序算法
平均时间复杂度为 O(n2)
*/
func main() {
	list := []int{2, 9, 4, 1, 3}
	InsertionSort(list)
	fmt.Printf("list : %v", list)
}

func InsertionSort(list []int) {
	length := len(list)
	if length <= 1 {
		return
	}
	for i := 1; i < length; i++ {
		//包留当前值
		v := list[i]
		j := i - 1
		for ; j >= 0; j-- {
			if list[j] > v {
				//如果大于则后移一位
				list[j+1] = list[j]
			} else {
				break
			}
		}
		//插入对应位置
		list[j+1] = v
	}
}
