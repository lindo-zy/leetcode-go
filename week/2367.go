package main

import "fmt"

func arithmeticTriplets(nums []int, diff int) int {
	n := len(nums)
	//转换为哈希表
	record := make(map[int]bool)
	for i := 0; i < n; i++ {
		record[nums[i]] = true
	}
	count := 0
	for _, num := range nums {
		if record[num-diff] && record[num+diff] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3)) //2
}
