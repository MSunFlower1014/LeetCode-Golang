package main

/*
141. 环形链表
给定一个链表，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 true 。 否则，返回 false 。



进阶：

你能用 O(1)（即，常量）内存解决此问题吗？


*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
哈希表判断节点是否访问过
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	m := make(map[*ListNode]bool)
	one := head
	for one != nil {
		if m[one] {
			return true
		}
		m[one] = true
		one = one.Next
	}
	return false
}

/*
快慢指针
慢指针起点为had，快指针起点为head.Next
慢指针一次走一步，快指针每次走两步
如果快慢指针相遇则存在环
*/
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	one, two := head, head
	for two != nil && two.Next != nil {
		if one == two {
			return true
		}
		one = one.Next
		two = two.Next.Next
	}
	return false
}
