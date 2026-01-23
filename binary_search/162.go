package main

import "fmt"

func findPeakElement(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
func main() {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
