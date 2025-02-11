package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {
	maxIndex := 0
	for i, num := range nums {
		if maxIndex >= i {
			maxIndex = max(maxIndex, i+num)
		}
	}
	return maxIndex >= len(nums)-1
}
