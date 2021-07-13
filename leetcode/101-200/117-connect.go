package main

/*
117. 填充每个节点的下一个右侧节点指针 II
给定一个二叉树

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。



进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
*/

func connect117(root *Node) *Node {
	if root == nil {
		return root
	}
	levelList := make([]*Node, 0)
	levelList = append(levelList, root)

	for len(levelList) > 0 {
		length := len(levelList)
		tempList := make([]*Node, 0)
		for i := 0; i < length; i++ {
			tempNode := levelList[i]
			if i == length-1 {
				tempNode.Next = nil
			} else {
				tempNode.Next = levelList[i+1]
			}

			if tempNode.Left != nil {
				tempList = append(tempList, tempNode.Left)
			}
			if tempNode.Right != nil {
				tempList = append(tempList, tempNode.Right)
			}
		}
		levelList = tempList
	}
	return root
}
