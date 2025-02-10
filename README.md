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

#### [剑指 Offer 06. 从尾到头打印链表](https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)

```
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint(head.Next), head.Val)

}
```

#### [剑指 Offer 58 - I. 翻转单词顺序](https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/)

```
func reverseWords(s string) string {
    s = strings.Trim(s, " ")
    words := strings.Fields(s)
    i, j := 0, len(words) - 1
    for i < j {
        words[i], words[j] = words[j], words[i]
        i++
        j--    
    }
    return strings.Join(words, " ")
}
```

```
[1, 2, 3, 4, 5, 10, 11, 15, 17, 19, 20, 21, 22, 23, 31, 32, 33, 34, 39, 42, 46, 48, 49, 53, 55, 56, 62, 64, 70, 72, 75, 76, 78, 79, 84, 85, 94, 96, 98, 101, 102, 104, 105, 114, 121, 122, 124, 128, 136, 139, 141, 142, 146, 148, 152, 155, 160, 169, 198, 200, 206, 207, 208, 215, 221, 226, 234, 236, 238, 239, 240, 279, 283, 287, 297, 300, 301, 309, 312, 322, 337, 338, 347, 394, 399, 406, 416, 437, 448, 461, 494, 538, 543, 560, 581, 621, 647, 617, 739, 72]
```