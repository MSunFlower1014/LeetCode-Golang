package main

import (
	"sort"
	"testing"
)

/*
sort.Sort() 方法
1.获取最大深度，值为 2*ceil(lg(n+1))
2.如果元素数量大于12，使用快排
3.如果元素数量小于12，使用希尔排序（改进版插入排序）
4.如果maxDepth为0，使用堆排序

快排选择区间首位，末位，中间位的中位数作为分区点
如果区间长度大于40，会八等份区间，将八点的中位数作为分区点

希尔排序通过一定的间隔gap将元素划分成几个区域来先进行排序，然后逐步缩小间隔进行排序，最后采用插入排序
此时选择的gap为6，因为最大长度为12
*/
func TestMaxDepth(t *testing.T) {

	result := maxDepth(1000)
	//sort.Sort(nil)

	t.Logf("100 num depth is %v", result)

	a := []int{1, 2, 3, 4, 5, 6}
	sort.Ints(a)
	t.Logf("a sorted is %v\n", a)
}

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}
