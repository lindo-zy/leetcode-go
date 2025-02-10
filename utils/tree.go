package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// 定义树节点
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 根据数组创建二叉树
func CreateTree(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	nodesStr := strings.Split(data, ",")
	if nodesStr[0] == "null" {
		return nil
	}

	root := &TreeNode{Val: parseInt(nodesStr[0])}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nodesStr) {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			node.Left = createNode(nodesStr, &i)
			node.Right = createNode(nodesStr, &i)
			queue = append(queue, node.Left, node.Right)
		}
	}
	return root
}

func createNode(nodes []string, index *int) *TreeNode {
	if *index >= len(nodes) || nodes[*index] == "null" {
		*index++
		return nil
	}
	node := &TreeNode{Val: parseInt(nodes[*index])}
	*index++
	return node
}

func parseInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func (root *TreeNode) PrintTree() {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	root.Left.PrintTree()
	root.Right.PrintTree()
}
