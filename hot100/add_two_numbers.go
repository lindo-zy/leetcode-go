package hot100

import (
	"leetcode-go/utils"
)

type ListNode = utils.ListNode

func SolutionAddTwoNumbers() {
	//构建链表
	l1 := utils.CreateLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := utils.CreateLinkedList([]int{9, 9, 9, 9})
	utils.PrintLinkedList(l1)
	utils.PrintLinkedList(l2)
	numbers := addTwoNumbers(l1, l2)
	utils.PrintLinkedList(numbers)
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
		sum := carry + x + y
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
