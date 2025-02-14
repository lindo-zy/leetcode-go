package main

import (
	"fmt"
	. "leetcode-go/utils"
)

func main() {
	fmt.Println(maxDepth(CreateTree("3,9,20,null,null,15,7")))

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
		ans = max(ans, level)

	}
	dfs(root, 1)
	return ans
}
