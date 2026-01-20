package main

import "fmt"

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
func main() {

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
