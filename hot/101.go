package main

import (
	"fmt"
	. "leetcode-go/utils"
)

func main() {
	nodes := "1,2,2,nil,3,nil,3"
	tree := CreateTree(nodes)
	res := isSymmetric(tree)
	fmt.Println(res)
}

func isSymmetric(root *TreeNode) bool {
	//左边和右边一样，即为对称
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}
func isMirror(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	l := isMirror(p.Left, q.Right)
	r := isMirror(p.Right, q.Left)
	return l && r && p.Val == q.Val
}
