package main

import "fmt"

func main() {
	fmt.Println(leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
}

func leastInterval(tasks []byte, n int) int {
	counter := make(map[byte]int)
	for _, task := range tasks {
		counter[task]++
	}
	maxFreq := 0
	var maxCount int
	for _, count := range counter {
		if count > maxFreq {
			maxFreq = count
			maxCount = 1
		} else if count == maxFreq {
			maxCount += 1
		}
	}
	minTime := (maxFreq-1)*(n+1) + maxCount
	if minTime > len(tasks) {
		return minTime
	}

	return len(tasks)
}
