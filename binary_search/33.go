package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target < nums[n-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}

		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
