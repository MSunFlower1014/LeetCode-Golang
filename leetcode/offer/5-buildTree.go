package main

/*
剑指 Offer 07. 重建二叉树
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。



例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

*/
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
3为根节点

寻找3在 inorder的下标为1
则 inorder[0:1] 为左子树元素
inorder[1+1:] 为右子树元素
inorder[0:1] 的长度为l1
则 preorder[1:l1+1] 为preorder中左子树元素
preorder[l1+1:] 为preorder中右子树元素
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[0:i])+1], inorder[0:i])
	root.Right = buildTree(preorder[len(inorder[0:i])+1:], inorder[i+1:])
	return root
}
