package main

import (
	"fmt"
	. "leetcode-go/utils"
)

func main() {
	nodes := "1,2,2,nil,3,nil,3"
	tree := CreateTree(nodes)
	fmt.Println(inorderTraversal(tree))

}

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return ans
}
