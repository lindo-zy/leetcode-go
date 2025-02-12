package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 4}, {3, 5}}))
}

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, interval := range intervals {
		if len(ans) > 0 && interval[0] <= ans[len(ans)-1][1] {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], interval[1])
		} else {
			ans = append(ans, interval)
		}
	}
	return ans
}
