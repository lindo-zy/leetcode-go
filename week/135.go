package main

import "fmt"

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
}
func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	record := make([]int, n)
	for i := 0; i < n; i++ {
		record[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			record[i] = record[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			record[i] = max(record[i], record[i+1]+1)
		}
	}
	total := 0
	for _, r := range record {
		total += r
	}
	return total
}
