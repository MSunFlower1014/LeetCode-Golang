package main

/*
剑指 Offer 04. 二维数组中的查找
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，
输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。
*/
func main() {

}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	height := len(matrix)
	if height == 0 {
		return false
	}
	width := len(matrix[0])
	if width == 0 {
		return false
	}

	i := 0
	for ; i < height && i < width && matrix[i][i] <= target; i++ {
		if matrix[i][i] == target {
			return true
		}
	}

	for k := 0; k < i; k++ {
		for l := i; l < width; l++ {
			if matrix[k][l] == target {
				return true
			}
		}
	}

	for k := i; k < height; k++ {
		for l := 0; l < i; l++ {
			if matrix[k][l] == target {
				return true
			}
		}
	}
	return false
}
