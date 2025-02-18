package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[1] = nums[0]
	for i := 2; i < n+1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n-1]
}
