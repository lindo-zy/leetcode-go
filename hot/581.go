package main

import "fmt"

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	left, right := -1, -2
	minRight, maxLeft := nums[n-1], nums[0]

	for i := 1; i < n; i++ {
		maxLeft = max(maxLeft, nums[i])
		minRight = min(minRight, nums[n-1-i])
		if nums[i] < maxLeft {
			right = i
		}
		if nums[n-1-i] > minRight {
			left = n - 1 - i
		}
	}

	return right - left + 1
}
