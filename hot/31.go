package main

import "fmt"

func main() {
	nums := []int{1, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	for i >= 0 && nums[i] >= nums[i+1] {
		i -= 1
	}
	if i < 0 {
		for j := 0; j < n/2; j++ {
			nums[j], nums[n-1-j] = nums[n-1-j], nums[j]
		}
		return
	}

	j := n - 1
	for j >= i && nums[j] <= nums[i] {
		j -= 1
	}
	nums[i], nums[j] = nums[j], nums[i]
	left, right := i+1, n-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

}
