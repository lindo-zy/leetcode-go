package main

import (
	"fmt"
	. "leetcode-go/utils"
	"slices"
)

func main() {
	fmt.Println(isPalindrome(CreateLinkedList([]int{1, 3, 2, 1})))
}

func isPalindrome(head *ListNode) bool {
	var nodes []int
	for head != nil {
		nodes = append(nodes, head.Val)
		head = head.Next
	}
	tmp := make([]int, len(nodes))
	copy(tmp, nodes)
	slices.Reverse(nodes)
	return slices.Equal(tmp, nodes)
}
