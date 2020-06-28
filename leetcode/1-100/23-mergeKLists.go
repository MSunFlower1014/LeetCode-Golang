package main

/*
23. 合并K个排序链表
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6

方法1：遍历所有链表，合并
方法2：分治算法，大大减少了合并次数
*/
func main() {

}

func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode = merge(lists, 0, len(lists)-1)
	return result
}

func merge(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (r + l) >> 1
	return mergeTwo(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = &ListNode{Val: -1}
	var prev = result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}
	return result.Next
}
