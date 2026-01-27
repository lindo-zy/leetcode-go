package main

import (
	"fmt"
	"slices"
)

func searchMatrix2(matrix [][]int, target int) bool {
	for _, row := range matrix {
		i, _ := slices.BinarySearch(row, target)
		if i < len(matrix[0]) && row[i] == target {
			return true
		}
	}

	return false
}
func main() {
	fmt.Println(searchMatrix2([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 13))
}
