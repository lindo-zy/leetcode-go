package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] = nums[i] + max(nums[i-1], 0)
	}
	return slices.Max(nums)
}
