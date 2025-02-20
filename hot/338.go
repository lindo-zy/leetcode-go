package main

import "fmt"

func main() {
	fmt.Println(countBits(5))
}
func countBits(n int) []int {
	return make([]int, n+1)
}
