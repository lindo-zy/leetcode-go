package main

import . "leetcode-go/utils"

func main() {
	list1 := CreateLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	list2 := CreateLinkedList([]int{9, 9, 9, 9})
	res := addTwoNumbers(list1, list2)
	PrintLinkedList(res)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum := x + y + carry
		if sum >= 10 {
			p.Next = &ListNode{Val: sum - 10}
			carry = 1
		} else {
			p.Next = &ListNode{Val: sum}
			carry = 0
		}
		p = p.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
