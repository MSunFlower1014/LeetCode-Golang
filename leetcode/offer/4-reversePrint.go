package main

/*
剑指 Offer 06. 从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
*/
func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var result []int
	newHead := reverseList(head)
	for newHead != nil {
		result = append(result, newHead.Val)
		newHead = newHead.Next
	}
	return result
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return result
}
