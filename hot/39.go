package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(candidates)

	var dfs func(nums []int, target int, tmp []int)
	dfs = func(nums []int, target int, tmp []int) {
		if target == 0 {
			combination := make([]int, len(tmp))
			copy(combination, tmp)
			ans = append(ans, combination)
			return
		}
		for i := 0; i < len(nums); i++ {
			if nums[i] <= target {
				dfs(nums[i:], target-nums[i], append(tmp, nums[i]))
			}
		}

	}
	dfs(candidates, target, make([]int, 0))
	return ans
}
