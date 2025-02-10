package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7} //49
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	n := len(height)
	ans := 0
	left := 0
	right := n - 1

	for left < right {
		ans = max(ans, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return ans
}
