package main

import "strings"

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	words := strings.Fields(s)
	i, j := 0, len(words)-1
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}
	return strings.Join(words, " ")
}

func main() {

	reverseWords("  hello world!  ")
}
