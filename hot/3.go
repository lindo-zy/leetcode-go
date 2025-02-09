package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aabaab!bb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	tmp := ""
	ans := 0
	for i := range s {
		if strings.Contains(tmp, string(s[i])) {
			index := min(strings.IndexByte(tmp, s[i])+1, len(tmp))
			tmp = tmp[index:]
		}
		tmp += string(s[i])
		ans = max(ans, len(tmp))
	}
	return ans
}
