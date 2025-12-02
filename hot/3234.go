package main

import "fmt"

func numberOfSubstrings(s string) int {
	//滑动窗口
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		zeros := 0
		ones := 0
		for j := i; j < n; j++ {
			if s[j] == '1' {
				ones++
			} else {
				zeros++
			}
		}
		if ones >= zeros*zeros {
			res++
		}
	}
	return res

}
func main() {
	fmt.Println(numberOfSubstrings("00011")) //错误代码
}
