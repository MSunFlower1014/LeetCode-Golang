package main

import "fmt"

/*
51. N皇后
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。



上图为 8 皇后问题的一种解法。

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

*/
func main() {
	fmt.Print(solveNQueens(4))
}

var rows []int
var hills []int
var dales []int
var num int

var queenResult [][]string
var queens []int

/*
判断是否 冲突
*/
func isNotUnderAttack(row, col int) bool {
	res := rows[col] + hills[row-col+2*num] + dales[row+col]
	return res == 0
}

func placeQueen(row, col int) {
	queens[row] = col
	rows[col] = 1
	hills[row-col+2*num] = 1
	dales[row+col] = 1
}

func removeQueen(row, col int) {
	queens[row] = 0
	rows[col] = 0
	hills[row-col+2*num] = 0
	dales[row+col] = 0
}

func addSolution() {
	var solution []string
	for i := 0; i < num; i++ {
		col := queens[i]
		s := ""
		for j := 0; j < col; j++ {
			s = s + "."
		}
		s = s + "Q"

		for j := 0; j < num-col-1; j++ {
			s = s + "."
		}
		solution = append(solution, s)
	}
	queenResult = append(queenResult, solution)
}

func queenBacktrack(row int) {
	for col := 0; col < num; col++ {
		if isNotUnderAttack(row, col) {
			placeQueen(row, col)

			if row+1 == num {
				addSolution()
			} else {
				queenBacktrack(row + 1)
			}

			removeQueen(row, col)
		}
	}
}

func solveNQueens(n int) [][]string {
	queenResult = [][]string{}
	num = n
	rows = make([]int, num)
	hills = make([]int, 4*num-1)
	dales = make([]int, 2*num-1)
	queens = make([]int, num)
	queenBacktrack(0)
	return queenResult
}
