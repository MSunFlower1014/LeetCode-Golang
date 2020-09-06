package main

import "fmt"

/*
62. 不同路径
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？


例如，上图是一个7 x 3 的网格。有多少可能的路径？


示例 1:

输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
*/
func main() {
	fmt.Print(uniquePaths2(7, 3))
}

/*
记录每个坐标可能到达的路线数
df [j][i] = df [j-1][i] + df[j][i-1]
单点路线数为左-1和上-1的路线和
df的第一行和首列 都为1
*/
func uniquePaths2(m int, n int) int {
	var df [][]int
	var temp []int
	for j := 0; j < m; j++ {
		temp = append(temp, 1)
	}
	df = append(df, temp)
	for i := 0; i < n; i++ {
		temp = make([]int, m)
		temp[0] = 1
		df = append(df, temp)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			df[j][i] = df[j-1][i] + df[j][i-1]
		}
	}

	return df[n-1][m-1]
}

/*
递归方法超出时间限制。。。
*/
var uniquePathsResult int = 0

func uniquePaths(m int, n int) int {
	uniquePathsResult = 0
	nextPath(m-1, n-1, 0, 0)
	return uniquePathsResult
}
func nextPath(m int, n int, x, y int) {
	if x > m || y > n {
		return
	}
	if x == m && y == n {
		uniquePathsResult++
		return
	}
	nextPath(m, n, x+1, y)
	nextPath(m, n, x, y+1)
}
