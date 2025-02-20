package main

import (
	"fmt"
	. "leetcode-go/utils"
)

func main() {

	fmt.Println(pathSum(CreateTree("5,4,8,11,null,13,4,7,2,null,null,5,1"), 22))

}

func pathSum(root *TreeNode, targetSum int) int {
	pre := make(map[int]int)
	pre[0] = 1
	var dfs func(root *TreeNode, cnt int) int
	dfs = func(root *TreeNode, cnt int) int {
		if root == nil {
			return 0
		}
		cnt += root.Val
		ans := pre[cnt-targetSum]
		pre[cnt] += 1
		if root.Left != nil {
			ans += dfs(root.Left, cnt)
		}
		if root.Right != nil {
			ans += dfs(root.Right, cnt)
		}
		pre[cnt] -= 1
		return ans
	}

	return dfs(root, 0)
}
