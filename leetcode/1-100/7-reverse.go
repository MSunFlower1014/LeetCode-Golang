package main

import (
	"fmt"
)

/*
7. 整数反转
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
*/
func main() {
	result := reverse(1534236469)
	fmt.Printf("%d", result)
}

func reverse(x int) int {
	const IntMax = int(^uint32(0) >> 1)
	const IntMin = ^IntMax
	var result int
	for x != 0 {
		pop := x % 10
		x = x / 10

		result = 10*result + pop
		if result > IntMax || result < IntMin {
			return 0
		}
	}
	return result
}
