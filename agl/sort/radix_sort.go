package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

func radixSort(arr []int) []int {
	size := len(strconv.Itoa(slices.Max(arr)))
	result := make([]int, len(arr))
	copy(result, arr)

	for i := 0; i < size; i++ {
		buckets := make([][]int, 10)
		divisor := int(math.Pow10(i))
		for _, num := range result {
			digit := (num / divisor) % 10
			buckets[digit] = append(buckets[digit], num)
		}
		idx := 0
		for j := 0; j < 10; j++ {
			for _, num := range buckets[j] {
				result[idx] = num
				idx++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(radixSort(radixSort([]int{6, 1, 5, 3, 4, 2})))
}
