package main

import (
	. "leetcode-go/utils"
)

func main() {
	PrintLinkedList(reverseList(CreateLinkedList([]int{1, 2})))
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
