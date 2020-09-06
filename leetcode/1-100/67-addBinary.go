package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
67. 二进制求和
给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。



示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/
func main() {
	fmt.Print(addBinary("100", "110010"))
}

func addBinary(a string, b string) string {
	aList := strings.Split(a, "")

	bList := strings.Split(b, "")

	aLength := len(a)
	bLength := len(b)
	aSub := 0
	bSub := 0
	temp := 0
	aLength--
	bLength--
	s := ""
	for true {
		if aLength < 0 || bLength < 0 {
			break
		}
		aSub, _ = strconv.Atoi(aList[aLength])
		bSub, _ = strconv.Atoi(bList[bLength])
		sum := aSub + bSub + temp
		if sum > 1 {
			temp = 1
			sum = sum % 2
			s = strconv.Itoa(sum) + s
		} else {
			temp = 0
			s = strconv.Itoa(sum) + s
		}
		aLength--
		bLength--
	}

	for i := aLength; i >= 0; i-- {
		aSub, _ = strconv.Atoi(aList[i])

		aSub = aSub + temp
		if aSub > 1 {
			temp = 1
			aSub = aSub % 2
			s = strconv.Itoa(aSub) + s
		} else {
			temp = 0
			s = strconv.Itoa(aSub) + s
		}
		aLength--
	}
	for i := bLength; i >= 0; i-- {
		aSub, _ = strconv.Atoi(bList[i])
		aSub = aSub + temp
		if aSub > 1 {
			temp = 1
			aSub = aSub % 2
			s = strconv.Itoa(aSub) + s
		} else {
			temp = 0
			s = strconv.Itoa(aSub) + s
		}
	}
	if temp == 1 {
		s = "1" + s
	}

	return s
}
