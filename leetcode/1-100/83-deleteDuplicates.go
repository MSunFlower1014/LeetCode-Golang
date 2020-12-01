package main

/**
83. 删除排序链表中的重复元素
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
*/
func main() {

}

/*
链表如果需要修改头节点，新增一个空节点为头节点更好操作
*/
func deleteDuplicates83(head *ListNode) *ListNode {
	result := &ListNode{Val: 0, Next: head}

	head = result
	lastVal := 0
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			lastVal = head.Next.Val
			for head.Next.Next != nil && head.Next.Next.Val == lastVal {
				head.Next.Next = head.Next.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return result.Next
}

func deleteDuplicates83_2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	result := head
	lastVal := 0
	for head.Next != nil {
		if head.Val == head.Next.Val {
			lastVal = head.Val
			for head.Next != nil && head.Next.Val == lastVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return result
}
