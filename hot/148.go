package main

import (
	. "leetcode-go/utils"
)

func main() {
	PrintLinkedList(sortList(CreateLinkedList([]int{-1, 5, 3, 4, 0})))
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	cur := slow.Next
	slow.Next = nil
	return mergeLinkedList(sortList(head), sortList(cur))
}

func mergeLinkedList(node1 *ListNode, node2 *ListNode) *ListNode {
	node := &ListNode{Val: 0}
	cur := node
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			cur.Next, node1 = node1, node1.Next
		} else {
			cur.Next, node2 = node2, node2.Next
		}
		cur = cur.Next
	}
	if node1 != nil {
		cur.Next, node1 = node1, node1.Next
	}
	if node2 != nil {
		cur.Next, node2 = node2, node2.Next
	}
	return node.Next
}
