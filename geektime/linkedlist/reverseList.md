#单链表反转两种实现

## 1.常规
```go
//包留前驱节点，后继节点 用于交换
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
```

## 2.递归
```go
// head.Next.Next = head 成环
// head.Next = nil 断开环
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
```