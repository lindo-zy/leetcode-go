package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	p, q := 1, 1
	for i := 0; i < len(nums)-1; i++ {
		p *= nums[i]
		res[i+1] = p
	}
	for i := len(nums) - 1; i > 0; i-- {
		q *= nums[i]
		res[i-1] *= q
	}
	return res
}
