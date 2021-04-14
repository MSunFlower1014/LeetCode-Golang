package linkedlist

import "testing"

type TowWayListNode struct {
	Val  string
	Next *TowWayListNode
	Pre  *TowWayListNode
}

/*
字符串回文判断
由于回文串最重要的就是对称，那么最重要的问题就是找到那个中心，通过快慢指针（一个每步走1，一个每步走2）
用快指针每步两格走，当他到达链表末端的时候，慢指针刚好到达中心
通过双向链表，从中心分别获取前驱结点和后继结点，依次比较即可
*/
func TestPalindrome(t *testing.T) {
	n1 := &TowWayListNode{Val: "a"}
	n2 := &TowWayListNode{Val: "b", Pre: n1}
	n3 := &TowWayListNode{Val: "b", Pre: n2}
	n4 := &TowWayListNode{Val: "a", Pre: n3}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	result := IsPalindrome(n1)
	t.Logf("result is %v", result)
}

func IsPalindrome(node *TowWayListNode) bool {
	two, one := node, node

	for two.Next != nil && two.Next.Next != nil {
		one = one.Next
		two = two.Next.Next
	}

	right := one
	if two.Next != nil {
		right = one.Next
	}

	for one != nil {
		if one.Val != right.Val {
			return false
		}
		one = one.Pre
		right = right.Next
	}

	return true
}
