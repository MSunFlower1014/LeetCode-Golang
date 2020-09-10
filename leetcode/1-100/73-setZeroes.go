package main

/*
73. 矩阵置零
给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。

示例 1:

输入:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
*/
func main() {

}

func setZeroes(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])

	x := []int{}
	y := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				x = append(x, i)
				y = append(y, j)
			}
		}
	}

	for _, i := range x {
		matrix[i] = make([]int, m)
	}

	for j := 0; j < n; j++ {
		for _, i := range y {
			matrix[j][i] = 0
		}
	}
}

func setZeroes2(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	a := map[int][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				a[i] = append(a[i], j)
			}
		}
	}
	for k, v := range a {
		for i := 0; i < n; i++ {
			matrix[k][i] = 0
		}
		for i := 0; i < len(v); i++ {
			for j := 0; j < m; j++ {
				matrix[j][v[i]] = 0
			}
		}
	}
}
