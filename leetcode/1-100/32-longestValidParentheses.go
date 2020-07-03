package main

import "fmt"

/*
32. 最长有效括号
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/
func main() {
	fmt.Print(longestValidParentheses("(()())"))
}

func longestValidParentheses(s string) int {
	result := 0
	length := len(s)
	dp := make([]int, length)
	for i := range s {
		if i == 0 {
			continue
		}
		if (string(s[i])) == ")" {
			if (string(s[i-1])) == "(" {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && string(s[i-dp[i-1]-1]) == "(" {
				if (i - dp[i-1]) >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}

			}

			if result < dp[i] {
				result = dp[i]
			}
		}
	}
	return result
}
