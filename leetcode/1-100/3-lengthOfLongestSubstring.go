package main

import "fmt"

/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

滑动窗口
*/
func main() {
	result := lengthOfLongestSubstring("aabc")
	fmt.Print(result)
}

func lengthOfLongestSubstring(s string) int {
	var list []int32
	var result = 0
	for _, v := range s {
		for i, s := range list {
			if v == s {
				if len(list) == 1 {
					list = []int32{}
					continue
				}
				//重复则切除重复前部分
				list = list[i+1:]
				continue
			}
		}
		list = append(list, v)
		if result < len(list) {
			result = len(list)
		}
	}
	return result
}
