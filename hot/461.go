package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 4))
}
func hammingDistance(x int, y int) int {
	distance := 0
	z := x ^ y
	for z != 0 {
		distance += z & 1
		z >>= 1
	}
	return distance
}
