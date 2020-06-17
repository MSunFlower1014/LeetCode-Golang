package main

/*
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。


*/
func main() {
	resultNode := ListNode{Val: 1}
	resultNode2 := ListNode{Val: 2}
	print(resultNode2.Next.Val)
	print(&(resultNode.Next))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return &ListNode{}
	}
	var l1Val, l2Val int
	if l1 == nil {
		l1Val = 0
	} else {
		l1Val = l1.Val
	}
	if l2 == nil {
		l2Val = 0
	} else {
		l2Val = l2.Val
	}
	val := l1Val + l2Val
	var n = 0
	if val > 9 {
		n = 1
		val = val - 10
	}
	resultNode := ListNode{Val: val}
	addTwoNumbers4Result(l1.Next, l2.Next, &resultNode, n)
	return &resultNode
}

func addTwoNumbers4Result(l1 *ListNode, l2 *ListNode, result *ListNode, n int) {
	if l1 == nil && l2 == nil {
		if n == 1 {
			nCode := ListNode{Val: 1}
			result.Next = &nCode
		}
		return
	}
	var l1Val, l2Val int
	if l1 == nil {
		l1 = &ListNode{Val: 0}
	}
	if l2 == nil {
		l2 = &ListNode{Val: 0}
	}
	l1Val = l1.Val
	l2Val = l2.Val
	val := l1Val + l2Val + n
	n = 0
	if val > 9 {
		n = 1
		val = val - 10
	}
	newCode := ListNode{Val: val}
	result.Next = &newCode
	addTwoNumbers4Result(l1.Next, l2.Next, &newCode, n)
}
