package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cur := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			cur = nums[i]
			count++
		} else if nums[i] == cur {
			count++
		} else {
			count--
		}
	}

	return cur
}
