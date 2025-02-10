package main

import "fmt"

func main() {
	s := "babad"
	res := longestPalindrome(s)
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	n := len(s)
	maxLen := 0
	start := 0

	var check func(left int, right int) int
	check = func(left int, right int) int {
		for left >= 0 && right < n && s[left] == s[right] {
			left -= 1
			right += 1
		}
		return right - left - 1
	}
	for i := 0; i < n; i++ {
		len1 := check(i, i)
		len2 := check(i, i+1)
		curLen := max(len1, len2)
		if curLen > maxLen {
			maxLen = curLen
			start = i - (curLen-1)/2
		}
	}

	return s[start : start+maxLen]
}
