package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reconstructQueue([][]int{
		{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
	}))
}
func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	for _, person := range people {
		if person[1] >= len(res) {
			res = append(res, person)
		} else {
			// 在指定位置插入元素
			res = append(res[:person[1]], append([][]int{person}, res[person[1]:]...)...)
		}
	}
	return res
}
