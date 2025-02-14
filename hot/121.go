package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
func maxProfit(prices []int) int {
	ans := 0
	minPrice := math.MaxInt
	for _, p := range prices {
		minPrice = min(minPrice, p)
		ans = max(ans, p-minPrice)
	}
	return ans
}
