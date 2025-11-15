package main

import "fmt"

func main() {
	fmt.Println(countOfSubstrings("aeioqq", 1))
}

func countOfSubstrings(word string, k int) int {
	n := len(word)
	vowel := map[int32]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	left := 0
	right := 0
	for right < n {
		curWin := word[left:right]
		if checkVowel(curWin, vowel) {

		}
		right++
	}
	ans := 0
	return ans
}
func checkVowel(word string, vowel map[int32]bool) bool {
	for _, w := range word {
		if _, ok := vowel[w]; ok {
			return true
		}
	}
	return false
}
