package main

/*
19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。
*/
func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	var temp *ListNode = &ListNode{Val: 0, Next: head}
	var first *ListNode = head
	for first.Next != nil {
		length++
		first = first.Next
	}
	first = temp
	length = length - n
	for length > 0 {
		length--
		first = first.Next
	}
	first.Next = first.Next.Next
	return temp.Next
}
