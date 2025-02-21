package main

import (
	. "leetcode-go/utils"
)

func main() {
	mergeTrees(CreateTree("1,3,2,5"), CreateTree("2,1,3,null,4,null,7")).PrintTree()
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
