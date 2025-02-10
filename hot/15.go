package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)

	for i := 0; i < n; i++ {
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			cur := nums[i] + nums[left] + nums[right]
			if cur == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left += 1
				right -= 1
				//去除重复的元素
				for left < right && nums[left] == nums[left-1] {
					left += 1
				}
				for left < right && nums[right] == nums[right+1] {
					right -= 1
				}

			} else if cur < 0 {
				left += 1
			} else {
				right -= 1
			}
		}
	}

	return ans

}
