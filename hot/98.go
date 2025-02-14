package main

import (
	"fmt"
	. "leetcode-go/utils"
)

func main() {
	nodes := "2,2,2"
	fmt.Println(isValidBST(CreateTree(nodes)))
}

func isValidBST(root *TreeNode) bool {
	ans := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	for i := 1; i < len(ans); i++ {
		if ans[i] > ans[i-1] {
			continue
		} else {
			return false
		}
	}
	return true
}
