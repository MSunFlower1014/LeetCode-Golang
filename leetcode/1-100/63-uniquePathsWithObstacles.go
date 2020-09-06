package main

import "fmt"

/*
63. 不同路径 II
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？



网格中的障碍物和空位置分别用 1 和 0 来表示。

说明：m 和 n 的值均不超过 100。

示例 1:

输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: 2
*/
func main() {
	var obstacleGrid [][]int
	obstacleGrid = append(obstacleGrid, []int{1})
	obstacleGrid = append(obstacleGrid, []int{0})
	fmt.Print(uniquePathsWithObstacles(obstacleGrid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height := len(obstacleGrid)
	if height == 0 {
		return 0
	}
	width := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	temp := 1
	for i := 0; i < width; i++ {
		if obstacleGrid[0][i] == 1 {
			obstacleGrid[0][i] = 0
			temp = 0
		} else {
			obstacleGrid[0][i] = temp
		}
	}
	temp = 1
	for i := 1; i < height; i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
			temp = 0
		} else {
			obstacleGrid[i][0] = temp
		}
	}

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
		}
	}

	return obstacleGrid[height-1][width-1]
}
