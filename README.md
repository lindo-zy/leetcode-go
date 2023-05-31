# leetcode-go

leetcode by golang

[TOC]

## 剑指offer系列

#### [剑指 Offer 05. 替换空格](https://leetcode.cn/problems/ti-huan-kong-ge-lcof/)

```go
package main

import (
	"fmt"
	"regexp"
)

func replaceSpace(s string) string {
	reg := regexp.MustCompile(`\s`)
	return reg.ReplaceAllString(s, "%20")
}

func main() {
	s := "We are happy."
	r := replaceSpace(s)
	fmt.Println(r)
}

```

#### [剑指 Offer 58 - II. 左旋转字符串](https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/)

```go
package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

func main() {

	s := "abcdefg"
	k := 2
	r := reverseLeftWords(s, k)
	fmt.Println(r)
}

```

#### [剑指 Offer 20. 表示数值的字符串](https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/)

```go
#参考的网上的题解

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
```

