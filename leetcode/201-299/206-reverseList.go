package main

/*
206. 反转链表
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
每步首先保存 后继节点，然后将当前节点的后继节点改为保存的前驱节点
*/
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = pre
		pre = curr
		curr = temp
	}
	return pre
}

/*
递归实现

head.Next.Next = head 成环
head.Next = nil 断开环

*/
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
