package main

import "fmt"

/*
54. 螺旋矩阵
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
*/
func main() {
	matrix := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	}
	fmt.Print(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	var result []int = []int{}

	height := len(matrix) - 1
	if height < 0 {
		return result
	}
	width := len(matrix[0]) - 1

	result = spiralOut(0, 0, width, height, result, matrix)
	return result
}

func spiralOut(xTop, yTop, xBottom, yBottom int, result []int, matrix [][]int) []int {
	fmt.Printf("%p ", result)

	if xTop == xBottom && yTop < yBottom {
		for i := yTop; i <= yBottom; i++ {
			result = append(result, matrix[i][xTop])
		}
		return result
	}

	if yTop == yBottom && xTop < xBottom {
		for i := xTop; i <= xBottom; i++ {
			result = append(result, matrix[yTop][i])
		}
		return result
	}
	if yTop == yBottom && xTop == xBottom {
		result = append(result, matrix[xTop][yTop])
		return result
	}
	if xTop > xBottom || yTop > yBottom {
		return result
	}

	for i := xTop; i < xBottom; i++ {
		result = append(result, matrix[yTop][i])
	}
	for i := yTop; i < yBottom; i++ {
		result = append(result, matrix[i][xBottom])
	}
	for i := xBottom; i > xTop; i-- {
		result = append(result, matrix[yBottom][i])
	}
	for i := yBottom; i > yTop; i-- {
		result = append(result, matrix[i][xTop])
	}

	xTop++
	yTop++
	xBottom--
	yBottom--
	return spiralOut(xTop, yTop, xBottom, yBottom, result, matrix)
}
