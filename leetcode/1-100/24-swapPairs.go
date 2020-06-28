package main

/*
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
*/
func main() {

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var result *ListNode
	result = head.Next
	swapTwo(result, head)
	return result
}

func swapTwo(head *ListNode, two *ListNode) {
	temp := head.Next
	head.Next = two
	two.Next = temp
	if temp != nil && temp.Next != nil {
		temp2 := temp.Next
		two.Next = temp2
		swapTwo(temp2, temp)
	}
}
