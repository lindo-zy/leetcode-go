package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordBreak("applepenapple", []string{"apple", "penapple"}))
}

func wordBreak(s string, wordDict []string) bool {
	var dfs func(s string) bool
	cache := make(map[string]bool)
	dfs = func(s string) bool {
		if val, exists := cache[s]; exists {
			return val // 如果缓存中有结果，直接返回
		}
		if len(s) == 0 {
			return true
		}
		res := false
		for _, w := range wordDict {
			if strings.HasPrefix(s, w) {
				res = res || dfs(s[len(w):])
			}
		}
		cache[s] = res
		return res
	}
	return dfs(s)
}
