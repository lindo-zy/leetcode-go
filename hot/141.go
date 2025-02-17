package main

import (
	"fmt"
	. "leetcode-go/utils"
)

func main() {
	fmt.Println(hasCycle(CreateLinkedList([]int{3, 2, 0, -4})))
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
