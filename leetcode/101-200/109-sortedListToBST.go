package main

/*
109. 有序链表转换二叉搜索树
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
*/

func sortedListToBST(head *ListNode) *TreeNode {
	list := make([]int, 0)

	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	return helpToBST(list)
}

func helpToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	mid := length / 2

	temp := &TreeNode{Val: nums[mid], Left: nil, Right: nil}

	temp.Left = helpToBST(nums[:mid])
	temp.Right = helpToBST(nums[mid+1:])

	return temp
}
