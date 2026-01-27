package main

import "fmt"

import (
	. "leetcode-go/utils"
)

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
func main() {
	fmt.Println(countNodes(CreateTree("1,2,3,4,5,6")))
}
