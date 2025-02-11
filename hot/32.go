package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

func longestValidParentheses(s string) int {
	ans := 0
	stack := []int{-1}

	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1] //弹出栈顶
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				cur := i - stack[len(stack)-1]
				ans = max(ans, cur)
			}
		}
	}
	return ans
}
