package main

import (
	"fmt"
	. "leetcode-go/utils"
	"sort"
)

func main() {
	nodes := "3,9,20,null,null,15,7"
	tree := CreateTree(nodes)
	tree.PrintTree()
	fmt.Println("")
	order := levelOrder(tree)
	fmt.Println(order)
}

func levelOrder(root *TreeNode) [][]int {
	res := make(map[int][]int)

	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	ans := make([][]int, 0)
	keys := make([]int, 0, len(res))
	for key := range res {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, k := range keys {
		ans = append(ans, res[k])
	}
	return ans
}
