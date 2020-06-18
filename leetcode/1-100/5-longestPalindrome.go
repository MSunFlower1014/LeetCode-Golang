package main

import (
	"fmt"
	"strings"
)

/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。


*/
func main() {
	result := longestPalindrome("cbbd")
	fmt.Printf(result)
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return strings.Split(s, "")[0]
		}
	}

	var start, end = 0, 0
	for i := range s {
		if i == (len(s) - 1) {
			continue
		}
		//以单个字符作为扩散点
		var right, left = i, i
		for j := 1; j <= i; j++ {
			if i+j > len(s)-1 {
				break
			}
			if s[i-j] == s[i+j] {
				left = i - j
				right = i + j
			} else {
				break
			}
		}
		if right-left > end-start {
			start = left
			end = right
		}

		//以两个相同字符作为扩散点 如 baab中的aa
		if s[i] == s[i+1] {
			left, right = i, i+1
			odd := 1
			for j := 1; j <= i; j++ {
				if i+j+odd > len(s)-1 {
					break
				}
				if s[i-j] == s[i+j+odd] {
					left = i - j
					right = i + j + odd
				} else {
					break
				}
			}
			if right-left > end-start {
				start = left
				end = right
			}
		}
	}
	return s[start : end+1]
}
