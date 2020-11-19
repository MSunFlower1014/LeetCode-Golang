package main

import "fmt"

/*
79. 单词搜索
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
*/
func main() {
	var board [][]byte = make([][]byte, 3)
	board[0] = []byte{'A', 'B', 'C', 'E'}
	board[1] = []byte{'S', 'F', 'C', 'S'}
	board[2] = []byte{'A', 'D', 'E', 'E'}
	result := exist(board, "ABCCED")
	fmt.Println(result)
}

type pair struct{ x, y int }

//获取上下左右 坐标
var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	//记录已访问元素
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		//不符合单词字母，直接返回
		if board[i][j] != word[k] {
			return false
		}
		//已经获取完整单词，返回true
		if k == len(word)-1 {
			return true
		}
		vis[i][j] = true
		defer func() {
			//回溯还源
			vis[i][j] = false
		}()
		//遍历上下左右，继续查找
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
