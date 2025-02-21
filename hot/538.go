package main

import (
	. "leetcode-go/utils"
)

func main() {
	bst := convertBST(CreateTree("3,2,4,1"))
	bst.PrintTree()
}

func convertBST(root *TreeNode) *TreeNode {
	pre := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		root.Val += pre
		pre = root.Val
		dfs(root.Left)
	}
	dfs(root)
	return root
}
