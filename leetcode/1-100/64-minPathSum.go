package main

/*
64. 最小路径和
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/
func main() {

}

func minPathSum(grid [][]int) int {
	height := len(grid)
	if height == 0 {
		return 0
	}
	width := len(grid[0])

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[0][j] = grid[0][j] + grid[0][j-1]
				continue
			}

			if j == 0 {
				grid[i][0] = grid[i][0] + grid[i-1][0]
				continue
			}

			if grid[i-1][j] < grid[i][j-1] {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			}
		}
	}
	return grid[height-1][width-1]
}
