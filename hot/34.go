package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{2, 2}, 2))
}

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	left := binarySearch(nums, target)
	if left == len(nums) || nums[left] != target {
		return ans
	}
	right := binarySearch(nums, target+1)
	ans[0] = left

	if right < 1 {
		ans[1] = right
	} else {
		ans[1] = right - 1
	}
	return ans
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
