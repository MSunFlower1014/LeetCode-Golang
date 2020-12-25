package main

/*
92. 反转链表 II
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/
func main() {

}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var dummy = &ListNode{Next: head, Val: 0}
	leftNode := dummy
	for i := 1; i < m; i++ {
		leftNode = leftNode.Next
	}

	currentNode := leftNode.Next
	var pre *ListNode
	for i := 0; i <= n-m; i++ {
		temp := currentNode.Next
		currentNode.Next = pre
		pre = currentNode
		currentNode = temp
	}

	leftNode.Next.Next = currentNode
	leftNode.Next = pre
	return dummy.Next
}
