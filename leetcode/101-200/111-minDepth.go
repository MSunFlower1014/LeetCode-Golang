package main

/*
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。



示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：2
示例 2：

输入：root = [2,null,3,null,4,null,5,null,6]
输出：5
*/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	firstList := make([]*TreeNode, 1)
	firstList[0] = root
	result := 1
	for len(firstList) > 0 {
		temp := make([]*TreeNode, 0)
		for _, v := range firstList {
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
			if v.Left == nil && v.Right == nil {
				return result
			}
		}

		firstList = temp
		result++
	}
	return result
}
