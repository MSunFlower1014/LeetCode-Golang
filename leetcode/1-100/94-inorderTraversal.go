package main

/*
94. 二叉树的中序遍历
给定一个二叉树的根节点 root ，返回它的 中序 遍历。
*/
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	return inorderHandle(result, root)
}

func inorderHandle(result []int, root *TreeNode) []int {
	if root == nil {
		return result
	}
	if root.Left != nil {
		result = inorderHandle(result, root.Left)
	}

	result = append(result, root.Val)

	if root.Right != nil {
		result = inorderHandle(result, root.Right)
	}
	return result
}
