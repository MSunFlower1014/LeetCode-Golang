package main

/*
100. 相同的树
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1:

输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true
示例 2:

输入:      1          1
          /           \
         2             2

        [1,2],     [1,null,2]

输出: false
*/
func main() {

}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (q == nil && p != nil) || (q != nil && p == nil) {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	if (p.Right == nil && q.Right != nil) || (p.Right != nil && q.Right == nil) {
		return false
	}

	if p.Right != nil && p.Right.Val != q.Right.Val {
		return false
	}

	if (p.Left == nil && q.Left != nil) || (p.Left != nil && q.Left == nil) {
		return false
	}

	if p.Left != nil && p.Left.Val != q.Left.Val {
		return false
	}

	if p.Left != nil {
		left := isSameTree(p.Left, q.Left)
		if !left {
			return false
		}
	}

	if p.Right != nil {
		right := isSameTree(p.Right, q.Right)
		if !right {
			return false
		}
	}
	return true
}
