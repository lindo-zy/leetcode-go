package main

import (
	"fmt"
	"slices"
	"sort"
)

func bucketSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	minNum := slices.Min(arr)
	maxNum := slices.Max(arr)

	buckets := make([][]int, n)
	rangePerBucket := float64(maxNum-minNum+1) / float64(n)
	for _, value := range arr {
		index := int(float64(maxNum-minNum) / rangePerBucket)
		if index >= n {
			index = n - 1
		}
		buckets[index] = append(buckets[index], value)
	}
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if len(buckets[i]) > 0 {
			sort.Ints(buckets[i])
			result = append(result, buckets[i]...)
		}
	}
	return result
}

func main() {
	fmt.Println(bucketSort([]int{5, 4, 1, 3, 2}))

}
