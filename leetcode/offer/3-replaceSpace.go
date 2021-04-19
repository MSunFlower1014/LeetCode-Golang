package main

import (
	"fmt"
	"strings"
)

/*
剑指 Offer 05. 替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."
*/
func main() {
	result := replaceSpace("We are happy.")
	fmt.Print(result)
}
func replaceSpace(s string) string {
	var result []string
	for _, v := range s {
		if string(v) == " " {
			result = append(result, "%20")
		} else {
			result = append(result, string(v))
		}
	}
	return strings.Join(result, "")
}
