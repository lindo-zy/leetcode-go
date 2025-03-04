package main

import "fmt"

func main() {
	fmt.Println(numTimesAllBlue([]int{3, 2, 4, 1, 5}))
}
func numTimesAllBlue(flips []int) int {
	ans := 0
	cm := 1
	for i, flip := range flips {
		cm = max(cm, flip)
		if 1+i == cm {
			ans += 1
		}
	}
	return ans
}
