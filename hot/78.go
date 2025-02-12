package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	var dfs func(start int, tmp []int)
	dfs = func(start int, tmp []int) {
		cur := make([]int, len(tmp))
		copy(cur, tmp)
		ans = append(ans, cur)
		for i := start; i < n; i++ {
			dfs(i+1, append(tmp, nums[i]))
		}
	}

	dfs(0, []int{})

	return ans
}
