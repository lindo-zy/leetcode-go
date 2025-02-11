package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	n := len(height)
	left := 0
	right := n - 1
	leftVal := 0
	rightVal := 0
	ans := 0
	for left < right {
		leftVal = max(leftVal, height[left])
		rightVal = max(rightVal, height[right])
		if height[left] < height[right] {
			ans += leftVal - height[left]
			left += 1
		} else {
			ans += rightVal - height[right]
			right -= 1
		}
	}
	return ans
}
