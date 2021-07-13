package main

/*
118. 杨辉三角
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

func generate(numRows int) [][]int {
	result := make([][]int, 0)

	firstLevel := make([]int, 1)
	firstLevel[0] = 1
	result = append(result, firstLevel)

	for i := 1; i < numRows; i++ {
		temp := make([]int, 0)
		temp = append(temp, 1)
		for j := 1; j < i; j++ {
			temp = append(temp, result[i-1][j]+result[i-1][j-1])
		}
		temp = append(temp, 1)
		result = append(result, temp)
	}

	return result
}
