package main

func main() {

}

/*
876. 链表的中间结点
给定一个头结点为 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
快慢指针
奇数长度：
1-2-3
当two.Next=nil（节点3）停止，one为中间节点

偶数长度：
1-2-3-4
当two=nil（节点4的后继节点）停止，one为中间节点
*/
func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	one, two := head, head

	for two != nil && two.Next != nil {
		one = one.Next
		two = two.Next.Next
	}
	return one
}
