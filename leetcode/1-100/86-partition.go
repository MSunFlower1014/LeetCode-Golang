package main

/*
86. 分隔链表
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。



示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
*/
func main() {

}

func partition(head *ListNode, x int) *ListNode {
	plusNode := &ListNode{}
	miniNode := &ListNode{}
	currentPlus := plusNode
	currentMini := miniNode
	for head != nil {
		if head.Val < x {
			currentMini.Next = head
			currentMini = head
		} else {
			currentPlus.Next = head
			currentPlus = head
		}
		head = head.Next
	}
	currentPlus.Next = nil
	currentMini.Next = plusNode.Next
	return miniNode.Next
}
