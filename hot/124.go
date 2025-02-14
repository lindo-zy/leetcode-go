package main

import (
	"fmt"
	"math"
)
import (
	. "leetcode-go/utils"
)

func main() {
	fmt.Println(maxPathSum(CreateTree("-10,9,20,null,null,15,7")))
}

func maxPathSum(root *TreeNode) int {
	ans := -math.MaxInt
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftValue := max(dfs(root.Left), 0)
		rightValue := max(dfs(root.Right), 0)
		ans = max(ans, leftValue+rightValue+root.Val)
		return root.Val + max(leftValue, rightValue)
	}
	dfs(root)
	return ans
}
