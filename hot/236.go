package main

import (
	. "leetcode-go/utils"
)

func main() {
	ancestor := lowestCommonAncestor(CreateTree("3,5,1,6,2,0,8,null,null,7,4"), CreateTree("5"), CreateTree("1"))
	ancestor.PrintTree()
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
