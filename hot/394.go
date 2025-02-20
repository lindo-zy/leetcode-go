package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(decodeString("abc3[cd]xyz"))
}
func decodeString(s string) string {
	res := ""
	num := ""
	st := make([][2]string, 0)
	for _, c := range s {
		if strings.Trim(string(c), "0123456789") == "" {
			num += string(c)
		} else if c == '[' {
			st = append(st, [2]string{res, num})
			res = ""
			num = ""
		} else if c == ']' {
			top := st[len(st)-1]
			st = st[:len(st)-1]
			atoi, _ := strconv.Atoi(top[1])
			res = top[0] + strings.Repeat(res, atoi)
		} else {
			res += string(c)
		}
	}

	return res
}
