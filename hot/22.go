package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	if n <= 0 {
		return ans
	}

	var backtrack func(left, right int, tmp string)
	backtrack = func(left, right int, tmp string) {
		if left > n || right > left {
			return
		}

		if len(tmp) == n*2 {
			ans = append(ans, tmp)
			return
		}

		backtrack(left+1, right, tmp+string('('))
		backtrack(left, right+1, tmp+string(')'))

	}
	backtrack(0, 0, "")

	return ans
}
