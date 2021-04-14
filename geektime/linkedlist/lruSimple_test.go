package linkedlist

import (
	"fmt"
	"strings"
	"testing"
)

type listNode struct {
	Val  int
	Next *listNode
}

type LinkedList struct {
	Head  *listNode
	Count int
	Cap   int
}

/*
LRU 缓存淘汰链表实现
维护一个有序单链表，越靠近链表尾部的结点是越早之前访问的。
当有一个新的数据被访问时，我们从链表头开始顺序遍历链表。
1. 如果此数据之前已经被缓存在链表中了，我们遍历得到这个数据对应的结点，并将其从原来的位置删除，然后再插入到链表的头部。
2. 如果此数据没有在缓存链表中，又可以分为两种情况：
如果此时缓存未满，则将此结点直接插入到链表的头部；
如果此时缓存已满，则链表尾结点删除，将新的数据结点插入链表的头部。
*/
func TestLinkedListLruSimple(t *testing.T) {
	linked := LinkedList{Cap: 10, Head: &listNode{}, Count: 0}
	for i := 0; i < 11; i++ {
		PutValue(&linked, i)
	}
	builder := new(strings.Builder)
	//10;9;8;7;6;5;4;3;2;1;
	t.Logf("linked : %v\n", GetLinkedString(builder, &linked))
	GetValue(&linked, 1)
	//1;10;9;8;7;6;5;4;3;2;
	t.Logf("linked : %v\n", GetLinkedString(builder, &linked))
	GetValue(&linked, 12)
	//12;1;10;9;8;7;6;5;4;3;
	t.Logf("linked : %v\n", GetLinkedString(builder, &linked))
}

func GetLinkedString(builder *strings.Builder, list *LinkedList) string {
	builder.Reset()
	head := list.Head
	for head != nil {
		builder.WriteString(fmt.Sprintf("%v;", head.Val))
		head = head.Next
	}
	return builder.String()
}

func PutValue(list *LinkedList, i int) {
	if list.Count == 0 {
		list.Head = &listNode{Val: i}
		list.Count++
		return
	}
	head := list.Head
	if head.Val == i {
		return
	}
	DeleteContainOrOut(list, i)
	head = list.Head
	first := &listNode{Val: i, Next: list.Head}
	list.Head = first
	list.Count++
}

func GetValue(list *LinkedList, key int) int {
	head := list.Head
	if head.Val == key {
		return key
	}
	DeleteContainOrOut(list, key)
	head = list.Head
	first := &listNode{Val: key, Next: list.Head}
	list.Head = first
	list.Count++
	return key
}

//1.已存在当前值，删除当前值（删除后会在链表头添加）
//2.超出容量，删除最后一个元素
func DeleteContainOrOut(list *LinkedList, key int) {
	head := list.Head
	temp := list.Head.Next
	for temp != nil {
		if temp.Val == key {
			head.Next = temp.Next
			list.Count--
			break
		}
		temp = temp.Next
		head = head.Next
	}

	if list.Count == list.Cap {
		head, temp = list.Head, list.Head.Next
		for temp.Next != nil {
			temp = temp.Next
			head = head.Next
		}
		head.Next = nil
		list.Count--
	}
}
