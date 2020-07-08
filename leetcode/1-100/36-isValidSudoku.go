package main

import "fmt"

/*
36. 有效的数独
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

数独部分空格内已填入了数字，空白格用 '.' 表示。
示例 1:

输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-sudoku
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	m := make(map[byte]bool)
	m['a'] = true
	fmt.Print(m['a'])
	m = make(map[byte]bool)
	fmt.Print(m['a'])
}

func isValidSudoku(board [][]byte) bool {
	m := make(map[byte]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if !m[board[i][j]] {
				m[board[i][j]] = true
			} else {
				return false
			}
		}
		m = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if !m[board[j][i]] {
				m[board[j][i]] = true
			} else {
				return false
			}
		}
		m = make(map[byte]bool)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !checkBox(board, i*3, j*3) {
				return false
			}
		}
	}
	return true
}

func checkBox(board [][]byte, x, y int) bool {
	m := make(map[byte]bool)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[x+i][y+j] == '.' {
				continue
			}
			if !m[board[x+i][y+j]] {
				m[board[x+i][y+j]] = true
			} else {
				return false
			}
		}
	}
	return true
}
