package main

import "fmt"

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minI := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minI] {
				minI = j
			}
		}
		if i != minI {
			arr[i], arr[minI] = arr[minI], arr[i]
		}
	}

	return arr
}

func main() {

	fmt.Println(selectionSort([]int{3, 2, 1, 5, 10, 4}))

}
