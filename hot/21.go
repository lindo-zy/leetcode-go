package main

import (
	. "leetcode-go/utils"
)

func main() {
	nodes1 := []int{1, 2, 4}
	nodes2 := []int{1, 3, 4}
	list1 := CreateLinkedList(nodes1)
	list2 := CreateLinkedList(nodes2)
	PrintLinkedList(mergeTwoLists(list1, list2))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return dummy.Next
}
