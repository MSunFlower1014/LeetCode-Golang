package runtime

import (
	"fmt"
	"testing"
)

/*
range表达式只会在for语句开始执行时被求值一次，无论后边会有多少次迭代；
range表达式的求值结果会被复制，也就是说，被迭代的对象是range表达式结果值的副本而不是原值。

数组的range中修改值不会影响 range 求值结果
切片为引用类型，复制后指向的数组仍为同一个，会影响 range 求值结果
*/
func TestRange(t *testing.T) {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			//当i=5时，e仍为6，因为求值为值复制
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	//[7 3 5 7 9 11]
	fmt.Println(numbers2)
}

/*
slice为引用类型
*/
func TestSliceRange(t *testing.T) {
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	numbers2 = numbers2[:]
	maxIndex2 := len(numbers2) - 1
	//切片操作和range操作的是同一对象
	for i, e := range numbers2 {
		if i == maxIndex2 {
			//当i=5时，e已经被改为21
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	//[22 3 6 10 15 21]
	fmt.Println(numbers2)
}
