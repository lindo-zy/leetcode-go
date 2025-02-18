package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	currentMax := nums[0]
	currentMin := nums[0]
	globalMax := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			currentMax, currentMin = currentMin, currentMax
		}
		currentMax = max(nums[i], currentMax*nums[i])
		currentMin = min(nums[i], currentMin*nums[i])
		globalMax = max(globalMax, currentMax)
	}
	return globalMax
}
