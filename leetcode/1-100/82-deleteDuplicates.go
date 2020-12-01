package main

/*
82. 删除排序链表中的重复元素 II
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
*/
func main() {

}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//创建一个空的头节点，方便获取链表头，省去如果头节点重复，需要重复修改头节点
	nilNode := &ListNode{Val: 0, Next: head}

	head = nilNode
	lastVal := 0
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			lastVal = head.Next.Val
			//存在重复的数字，遍历全部删除
			for head.Next != nil && head.Next.Val == lastVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return nilNode.Next
}
