package main

import (
	"fmt"
	"strings"
)

/*
6. Z 字形变换
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
*/
func main() {
	result := convert("LEETCODEISHIRING", 3)
	fmt.Printf(result)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var list []string
	for i := 0; i < numRows; i++ {
		list = append(list, "")
	}
	data := strings.Split(s, "")
	n := 0
	flag := true
	for _, v := range data {
		list[n] = list[n] + v
		if flag {
			n++
			if n == numRows-1 {
				flag = false
			}
		} else {
			n--
			if n == 0 {
				flag = true
			}
		}
	}
	result := ""
	for _, v := range list {
		result = result + v
	}
	return result
}
