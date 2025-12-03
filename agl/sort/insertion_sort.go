package main

import "fmt"

func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j -= 1
		}
		arr[j] = temp
	}
	return arr
}

func main() {
	fmt.Println(InsertionSort([]int{1, 2, 3, 10, 6, 5, 4}))
}
