package main

/*
107. 二叉树的层序遍历 II
给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层序遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	helpArray := make([][]*TreeNode, 0)

	firstLevel := make([]*TreeNode, 1)
	firstLevel[0] = root

	helpArray = append(helpArray, firstLevel)

	for len(firstLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		for _, v := range firstLevel {
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}

			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}
		}
		if len(nextLevel) == 0 {
			break
		}
		helpArray = append(helpArray, nextLevel)
		firstLevel = nextLevel
	}

	for i := len(helpArray) - 1; i >= 0; i-- {
		var temp []int
		for _, v := range helpArray[i] {
			temp = append(temp, v.Val)
		}
		result = append(result, temp)
	}
	return result
}
