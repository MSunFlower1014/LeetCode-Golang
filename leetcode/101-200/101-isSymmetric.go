package main

/*
101. 对称二叉树
给定一个二叉树，检查它是否是镜像对称的。



例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3


但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
*/
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return checkSymmetric(root, root)
}

func checkSymmetric(p, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}

	if q == nil || p == nil {
		return false
	}

	return q.Val == p.Val && checkSymmetric(q.Right, p.Left) && checkSymmetric(q.Left, p.Right)
}
