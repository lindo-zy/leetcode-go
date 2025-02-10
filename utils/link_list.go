package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateLinkedList 从切片构建链表
func CreateLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	// 创建头节点
	head := &ListNode{Val: nums[0]}
	current := head

	// 遍历切片，构建链表
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}

// PrintLinkedList 打印链表
func PrintLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
