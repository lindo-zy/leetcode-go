package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	ans := ""
	findout := make(map[byte]int)
	for i := range t {
		findout[t[i]] += 1
	}
	minLen := math.MaxInt
	n := len(s)
	counter := len(t)
	left := 0
	right := 0
	for right < n {
		if findout[s[right]] > 0 {
			counter -= 1
		}
		findout[s[right]] -= 1
		right++
		for counter == 0 {
			if minLen > right-left {
				minLen = right - left
				ans = s[left:right]
			}
			if findout[s[left]] == 0 {
				counter += 1
			}
			findout[s[left]] += 1
			left++
		}
	}

	return ans
}
