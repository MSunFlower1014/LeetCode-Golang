package main

import "math"

/*
98. 验证二叉搜索树
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
*/
func main() {

}

func isValidBST(root *TreeNode) bool {
	return checkBST(root, math.MinInt64, math.MaxInt64)
}

func checkBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	left := root.Left
	right := root.Right
	val := root.Val

	if val <= min || val >= max {
		return false
	}
	return checkBST(left, min, val) && checkBST(right, val, max)
}
