package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		index := abs(num) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}
	var res []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
