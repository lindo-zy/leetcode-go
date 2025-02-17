package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

func longestConsecutive(nums []int) int {
	ans := 0
	set := make(map[int]bool)
	for _, n := range nums {
		set[n] = true
	}
	for num := range set {
		if !set[num-1] {
			tmp := 1
			for set[num+1] {
				num += 1
				tmp += 1
			}
			ans = max(ans, tmp)
		}
	}
	return ans
}
