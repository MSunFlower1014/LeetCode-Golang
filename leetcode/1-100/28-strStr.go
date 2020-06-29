package main

import "fmt"

/*
28. 实现 strStr()
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1

"mississippi"
"issip"

"mississippi"
"pi"
*/
func main() {
	fmt.Print(strStr("mississippi", "pi"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	k := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[k] {
			if k == len(needle)-1 {
				return i - k
			}
			k++
		} else {
			if k != 0 {
				i = i - k + 1
			}
			if needle[0] == haystack[i] {
				k = 1

			} else {
				k = 0
			}
		}
	}
	return -1
}
