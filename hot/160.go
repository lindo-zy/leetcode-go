package main

import (
	. "leetcode-go/utils"
)

func main() {
	PrintLinkedList(getIntersectionNode(CreateLinkedList([]int{1, 9, 1, 2, 4}), CreateLinkedList([]int{3, 2, 4})))
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A, B := headA, headB
	for A != B {
		if A != nil {
			A = A.Next
		} else {
			A = headB
		}
		if B != nil {
			B = B.Next
		} else {
			B = headA
		}
	}
	return A
}
