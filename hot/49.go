package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	record := map[string][]string{}
	for _, str := range strs {
		key := sortString(str)
		record[key] = append(record[key], str)
	}
	for _, group := range record {
		ans = append(ans, group)
	}

	return ans
}
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
