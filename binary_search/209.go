package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left := 0
	right := 0
	ans := math.MaxInt32
	total := 0
	for right < n {
		total += nums[right]
		for total >= target {
			ans = min(ans, right-left+1)
			total -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
