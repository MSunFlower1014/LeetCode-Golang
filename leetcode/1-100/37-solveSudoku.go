package main

/*
37. 解数独
编写一个程序，通过已填充的空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。


*/
func main() {

}

var boardList = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func solveSudoku(board [][]byte) {
	setBoard(board, 0, 0)
}

func setBoard(board [][]byte, x, y int) bool {
	if x == 9 {
		return true
	}
	var n, m int
	if y >= 8 {
		n = x + 1
		m = 0
	} else {
		n = x
		m = y + 1
	}
	if board[x][y] != '.' {
		return setBoard(board, n, m)
	}

	for i := 0; i < 9; i++ {
		if checkBoard(board, x, y, boardList[i]) {
			board[x][y] = boardList[i]
			if setBoard(board, n, m) {
				return true
			} else {
				board[x][y] = '.'
			}
		}
	}
	board[x][y] = '.'
	return false
}

func checkBoard(board [][]byte, x, y int, n byte) bool {
	for i := 0; i < 9; i++ {
		if board[x][i] == n {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][y] == n {
			return false
		}
	}

	m := x / 3 * 3
	u := y / 3 * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[m+i][u+j] == n {
				return false
			}
		}
	}
	return true
}
