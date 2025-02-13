package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximalRectangle([][]byte{
		{'1'},
	}))
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	pre := make([]int, n+1)
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				pre[j] = pre[j] + 1
			} else {
				pre[j] = 0
			}
		}
		stack := []int{-1}
		for k, num := range pre {
			for len(stack) > 1 && pre[stack[len(stack)-1]] > num {
				index := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				ans = max(ans, pre[index]*(k-stack[len(stack)-1]-1))
			}
			stack = append(stack, k)
		}
	}

	return ans
}
