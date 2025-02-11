package main

import (
	. "leetcode-go/utils"
)

func main() {
	nodes := []int{1, 2, 3, 4, 5}
	list := CreateLinkedList(nodes)
	PrintLinkedList(removeNthFromEnd(list, 2))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//快慢指针
	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}
