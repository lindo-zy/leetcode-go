package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	a := 0
	for _, num := range nums {
		a ^= num
	}
	return a
}
