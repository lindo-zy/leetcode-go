package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := range nums {
		if _, ok := numMap[target-nums[i]]; ok {
			return []int{numMap[target-nums[i]], i}
		}
		numMap[nums[i]] = i
	}
	return nil
}
