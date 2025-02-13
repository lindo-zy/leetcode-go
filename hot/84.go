package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{1}))
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	ans := 0
	stack := []int{-1}
	for k, height := range heights {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > height {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = max(ans, heights[index]*(k-stack[len(stack)-1]-1))
		}
		stack = append(stack, k)
	}

	return ans
}
