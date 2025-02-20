package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 20))

}
func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		j, _ := slices.BinarySearch(row, target)
		if j < len(matrix[0]) && row[j] == target {
			return true
		}
	}
	return false
}
