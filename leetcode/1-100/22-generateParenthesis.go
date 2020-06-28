package main

import "fmt"

/*
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
*/
func main() {
	fmt.Print(generateParenthesis(5))
}

var caches = make(map[int][]string)

func generateParenthesis(n int) []string {
	return generate(n)
}

func generate(n int) []string {
	if caches[n] != nil {
		return caches[n]
	}

	var ans []string
	if n == 0 {
		ans = append(ans, "")
	} else {
		for c := 0; c < n; c++ {
			for _, left := range generate(c) {
				for _, right := range generate(n - 1 - c) {
					ans = append(ans, "("+left+")"+right)
				}
			}
		}
	}
	caches[n] = ans
	return ans
}
