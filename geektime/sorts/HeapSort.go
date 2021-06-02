package main

import "fmt"

/**
堆排序
下标为i的父节点下标为 (i-1)/2
下标为i的左孩子节点下标为 i*2 +1
下标为i的右孩子节点下标为 i*2 +2
建堆复杂度为 On

Heapify 复杂度为Ologn

总体复杂度为 N * logn

不稳定算法
*/
func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	length := len(nums)

	//从第一个非叶子节点开始 维护堆 即 长度/2-1
	for i := length/2 - 1; i >= 0; i-- {
		Heapify(nums, length, i)
	}
	fmt.Printf("heapify is %v \n", nums)

	for i := length - 1; i > 0; i-- {
		Swap(nums, 0, i)
		Heapify(nums, i, 0)
	}
	fmt.Printf("heapify is %v \n", nums)
}

func Heapify(nums []int, n, i int) {
	largest := i
	lSon := i*2 + 1
	rSon := i*2 + 2

	if lSon < n && nums[largest] < nums[lSon] {
		largest = lSon
	}

	if rSon < n && nums[largest] < nums[rSon] {
		largest = rSon
	}

	if largest != i {
		Swap(nums, i, largest)
		Heapify(nums, n, largest)
	}
}

func Swap(nums []int, s, n int) {
	temp := nums[n]
	nums[n] = nums[s]
	nums[s] = temp
}
