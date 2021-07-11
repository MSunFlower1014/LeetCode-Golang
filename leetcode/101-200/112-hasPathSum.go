package main

import "fmt"

/*
112. 路径总和
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。

叶子节点 是指没有子节点的节点。



示例 1：


输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
*/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil || targetSum < root.Val {
		return false
	}
	root.Val = targetSum - root.Val
	levelList := make([]*TreeNode, 1)
	levelList[0] = root

	for len(levelList) > 0 {
		tempList := make([]*TreeNode, 0)
		for _, v := range levelList {
			tempVal := v.Val
			if v.Val == 0 && v.Right == nil && v.Left == nil {
				return true
			}
			if v.Left != nil {
				v.Left.Val = tempVal - v.Left.Val
				tempList = append(tempList, v.Left)
			}

			if v.Right != nil {
				v.Right.Val = tempVal - v.Right.Val
				tempList = append(tempList, v.Right)
			}
		}
		levelList = tempList
		fmt.Printf("%s\n", levelList)
	}
	return false
}
