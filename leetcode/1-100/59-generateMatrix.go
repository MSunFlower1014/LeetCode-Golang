package main

import "fmt"

/*
59. 螺旋矩阵 II
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/
func main() {
	fmt.Print(generateMatrix(1))
}

func generateMatrix(n int) [][]int {
	var result [][]int
	if n == 0 {
		return result
	}
	for i := 0; i < n; i++ {
		result = append(result, make([]int, n))
	}
	return setNumber(result, 0, n-1, 1)
}

func setNumber(result [][]int, x, y, n int) [][]int {
	if x == y {
		result[x][x] = n
		return result
	}
	if x > y {
		return result
	}

	for i := x; i < y; i++ {
		result[x][i] = n
		n++
	}

	for i := x; i < y; i++ {
		result[i][y] = n
		n++
	}

	for i := y; i > x; i-- {
		result[y][i] = n
		n++
	}

	for i := y; i > x; i-- {
		result[i][x] = n
		n++
	}
	x++
	y--
	return setNumber(result, x, y, n)
}
