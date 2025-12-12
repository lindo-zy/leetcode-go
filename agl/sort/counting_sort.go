package main

import (
	"fmt"
	"slices"
)

func countingSort(arr []int) []int {
	arrMin := slices.Min(arr)
	arrMax := slices.Max(arr)
	size := arrMax - arrMin + 1
	counts := make([]int, size)
	for _, num := range arr {
		counts[num-arrMin] += 1
	}
	for i := 1; i < size; i++ {
		counts[i] += counts[i-1]
	}
	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		result[counts[num-arrMin]-1] = num
		counts[num-arrMin] -= 1
	}

	return result
}

func main() {
	fmt.Println(countingSort([]int{5, 4, 3, 2, 1}))
}
