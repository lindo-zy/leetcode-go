package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if n == 0 || m == 0 {
		return false
	}

	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2
		row := mid / n
		col := mid % n
		value := matrix[row][col]

		if value == target {
			return true
		} else if value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60}}, 3))
}
