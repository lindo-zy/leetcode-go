package main

import "fmt"

func ShellSort(arr []int) []int {
	n := len(arr)
	gap := n / 2

	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for ; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = temp
		}
		gap = gap / 2
	}

	return arr
}

func main() {
	fmt.Println(ShellSort([]int{3, 2, 1, 10, 5, 6, 4}))
}
