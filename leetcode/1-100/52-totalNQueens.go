package main

import "fmt"

/*
52. N皇后 II
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。



上图为 8 皇后问题的一种解法。

给定一个整数 n，返回 n 皇后不同的解决方案的数量。
*/
func main() {
	fmt.Print(totalNQueens(4))
}

var totalNum int
var Rows []int
var Hills []int
var Dales []int
var Num int
var Queens []int

func totalNQueens(n int) int {
	totalNum = 0
	Num = n
	Rows = make([]int, Num)
	Hills = make([]int, 4*Num-1)
	Dales = make([]int, 2*Num-1)
	Queens = make([]int, Num)
	totalNQueenBacktrack(0)
	return totalNum
}

func totalNQueenBacktrack(row int) {
	for col := 0; col < Num; col++ {
		if totalIsNotUnderAttack(row, col) {
			totalPlaceQueen(row, col)

			if row+1 == Num {
				totalNum++
			} else {
				totalNQueenBacktrack(row + 1)
			}

			totalRemoveQueen(row, col)
		}
	}
}

/*
判断是否 冲突
*/
func totalIsNotUnderAttack(row, col int) bool {
	res := Rows[col] + Hills[row-col+2*Num] + Dales[row+col]
	return res == 0
}

func totalPlaceQueen(row, col int) {
	Queens[row] = col
	Rows[col] = 1
	Hills[row-col+2*Num] = 1
	Dales[row+col] = 1
}

func totalRemoveQueen(row, col int) {
	Queens[row] = 0
	Rows[col] = 0
	Hills[row-col+2*Num] = 0
	Dales[row+col] = 0
}
