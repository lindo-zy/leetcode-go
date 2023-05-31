package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func isNumber2(s string) bool {
	//前后空格无效
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if len(s) == 0 {
		return false
	}

	if strings.Contains(s, "+-") {
		return false
	}

	if strings.Contains(s, "++") {
		return false
	}
	if strings.Contains(s, "--") {
		return false
	}
	if strings.Contains(s, "-+") {
		return false
	}
	if strings.Contains(s, "..") {
		return false
	}
	if strings.Contains(s, ".+") {
		return false
	}
	if strings.Contains(s, ".-") {
		return false
	}

	if s[0] == 'e' || s[0] == 'E' {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	if strings.Contains(s, ".e") || strings.Contains(s, "e.") {
		return false
	}

	//只能有一个e或者E,且后面必须是整数
	cntE := strings.Count(s, "e")
	if cntE > 1 {
		return false
	} else if cntE == 1 {
		//e后面不能有.号
		indexE := strings.Index(s, "e")
		if strings.Contains(s[indexE:], ".") {
			return false
		}
		//后面必须有数
		if indexE == len(s)-1 {
			return false
		}
		if strings.Index(s, "e+") == len(s)-2 {
			return false
		}

		if strings.Index(s, "e-") == len(s)-2 {
			return false
		}
	}

	if strings.Count(s, ".") > 1 {
		return false
	}

	s = strings.ReplaceAll(s, "e+", "")
	s = strings.ReplaceAll(s, "e-", "")
	s = strings.ReplaceAll(s, "e", "")
	s = strings.ReplaceAll(s, ".", "")

	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}
func isNumber(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	if strings.ContainsAny(s, " bopx") {
		return false
	}
	if i := strings.IndexAny(s, "e"); i >= 0 {
		n1, _ := fmt.Sscan(s[:i], new(big.Float), new(string))
		n2, _ := fmt.Sscanf(s[i+1:], "%d%s", new(big.Int), new(string)) // %d 强制以十进制方式读入
		return n1 == 1 && n2 == 1
	}
	n, _ := fmt.Sscan(s, new(big.Float), new(string))
	return n == 1
}

func main() {
	strList := []string{"0"}
	for i := range strList {
		r := isNumber(strList[i])
		fmt.Println(r)
	}

}
