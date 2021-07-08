package main

/*
110. 平衡二叉树
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。



示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：true
*/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	rootVal := root.Val
	left := root.Left
	right := root.Right

	result := true
	if left != nil {
		leftVal := left.Val
		if rootVal < leftVal {
			return false
		} else {
			result = result && isBalanced(left)
		}
	}

	if right != nil {
		rightVal := right.Val
		if rootVal < rightVal {
			return false
		} else {
			result = result && isBalanced(left)
		}
	}
	return result
}
