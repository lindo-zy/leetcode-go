package main

import (
	"fmt"
)

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var ans []string
	var path string
	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(digits) {
			ans = append(ans, path)
			return
		}
		for _, letter := range phoneMap[digits[index]] {
			path += string(letter)
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}
