package main

import "fmt"

/*
72. 编辑距离
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')


官方解答：https://leetcode-cn.com/problems/edit-distance/solution/bian-ji-ju-chi-by-leetcode-solution/
*/
func main() {
	fmt.Print(minDistance("horse", "ros"))
}

func minDistance(word1, word2 string) int {
	n := len(word1)
	m := len(word2)

	if n == 0 || m == 0 {
		return n + m
	}

	D := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		D[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		D[i][0] = i
	}

	for i := 0; i <= m; i++ {
		D[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			left := D[i-1][j] + 1
			down := D[i][j-1] + 1
			leftDown := D[i-1][j-1]

			if word1[i-1] != word2[j-1] {
				leftDown++
			}

			D[i][j] = getMin(left, getMin(down, leftDown))
		}
	}
	return D[n][m]
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
