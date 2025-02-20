package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))

}

func lengthOfLIS(nums []int) int {
	st := make([]int, 0)
	for _, num := range nums {
		idx, _ := slices.BinarySearch(st, num)
		if idx == len(st) {
			st = append(st, num)
		} else {
			st[idx] = num
		}
	}
	return len(st)
}
