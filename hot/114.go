package main

import (
	. "leetcode-go/utils"
)

func main() {
	flatten(CreateTree("1,2,5,3,4,null,6"))
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	tmp := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = tmp
}
