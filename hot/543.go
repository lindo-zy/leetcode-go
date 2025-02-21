package main

import "fmt"
import (
	. "leetcode-go/utils"
)

func main() {
	fmt.Println(diameterOfBinaryTree(CreateTree("1,2,3,4,5")))
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		ans = max(ans, left+right)
		return max(left, right) + 1
	}
	dfs(root)
	return ans
}
