package main

import "fmt"

/*
剑指 Offer 13. 机器人的运动范围
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，
因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？



示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 2：

输入：m = 3, n = 1, k = 0
输出：1
*/
func main() {
	fmt.Print(movingCount(38, 15, 9))
}

func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}

	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	ans := 1
	vis[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || SumIntDigits(i)+SumIntDigits(j) > k {
				continue
			}

			if i-1 >= 0 {
				vis[i][j] = UnionBool(vis[i-1][j], vis[i][j])
			}
			if j-1 >= 0 {
				vis[i][j] = UnionBool(vis[i][j-1], vis[i][j])
			}
			if vis[i][j] {
				ans++
			}
		}
	}
	return ans
}

func SumIntDigits(i int) int {
	result := 0
	for i != 0 {
		result = result + i%10
		i = i / 10
	}
	return result
}

func UnionBool(x, y bool) bool {
	if x {
		return true
	}
	if y {
		return true
	}
	return false
}
