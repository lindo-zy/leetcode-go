package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var res [][]int // 结果集
	n := len(nums)
	var dfs func([]int, []int)
	dfs = func(nums []int, tmp []int) {
		if len(tmp) == n {
			combination := make([]int, len(tmp))
			copy(combination, tmp)
			res = append(res, combination)
			return
		}
		for i := 0; i < len(nums); i++ {
			newNums := append(append([]int{}, nums[:i]...), nums[i+1:]...)
			dfs(newNums, append(tmp, nums[i]))
		}
	}

	dfs(nums, []int{})
	return res
}
