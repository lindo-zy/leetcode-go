package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	n := len(nums)
	left, right := 0, n-1
	i := 0
	for i <= right {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left += 1
			i++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right -= 1
		} else {
			i += 1
		}
	}
}
