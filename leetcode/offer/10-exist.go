package main

/*
剑指 Offer 12. 矩阵中的路径
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

var height, width int

func exist(board [][]byte, word string) bool {

	height = len(board)
	if height == 0 {
		return false
	}
	width = len(board[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == word[0] && checkBoard(word, board, i, j) {
				return true
			}
		}
	}
	return false
}

func checkBoard(word string, board [][]byte, i, j int) bool {
	if len(word) == 0 {
		return true
	}
	if i >= height || j >= width || i < 0 || j < 0 {
		return false
	}
	if board[i][j] == word[0] {
		board[i][j] = byte(0)
		if checkBoard(word[1:], board, i, j+1) || checkBoard(word[1:], board, i+1, j) ||
			checkBoard(word[1:], board, i-1, j) || checkBoard(word[1:], board, i, j-1) {
			return true
		} else {
			board[i][j] = word[0]
			return false
		}
	} else {
		return false
	}
}
