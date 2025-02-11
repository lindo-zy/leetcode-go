package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "([])"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	for strings.Contains(s, "()") || strings.Contains(s, "[]") || strings.Contains(s, "{}") {
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")
	}

	return len(s) == 0
}
