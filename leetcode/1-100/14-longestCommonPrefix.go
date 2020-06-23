package main

import "fmt"

/*
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/
func main() {
	strs := []string{"abc", "ab", "abdc"}
	fmt.Print(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	i := len(result)
	for _, v := range strs {
		if result == v {
			continue
		}
		if i > len(v) {
			i = len(v)
		}
		for ; i > 0; i-- {
			if result[0:i] == v[0:i] {
				result = result[0:i]
				break
			}
		}
		if i == 0 {
			result = ""
			break
		}
	}
	return result
}
