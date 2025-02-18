package main

import . "leetcode-go/utils"

func main() {
	tree := CreateTree("2,3,1")
	node := invertTree(tree)
	node.PrintTree()
}

func invertTree(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		cur := &TreeNode{Val: root.Val}
		cur.Left = dfs(root.Right)
		cur.Right = dfs(root.Left)
		return cur
	}
	return dfs(root)
}
