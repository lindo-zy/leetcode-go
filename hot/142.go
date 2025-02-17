package main

import (
	. "leetcode-go/utils"
)

func main() {
	PrintLinkedList(detectCycle(CreateLinkedList([]int{3, 2, 0, -4})))
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	record := make(map[*ListNode]bool)
	cur := head
	for cur != nil && cur.Next != nil {
		if _, ok := record[cur]; ok {
			return cur
		}
		record[cur] = true
		cur = cur.Next
	}
	return nil
}
