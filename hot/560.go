package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
func subarraySum(nums []int, k int) int {
	ans := 0
	n := len(nums)
	pre := make(map[int]int)
	pre[0] = 1
	s := 0
	for i := 0; i < n; i++ {
		s += nums[i]
		ans += pre[s-k]
		pre[s] += 1
	}
	return ans
}
