package main

import (
	"math/rand"
	"sort"
	"time"
)

type Solution struct {
	Sum    int
	Prefix []int
}

func constructor(w []int) Solution {
	rand.Seed(time.Now().UnixNano())
	var s Solution
	for _, v := range w {
		s.Sum += v
		s.Prefix = append(s.Prefix, s.Sum-1)
	}
	return s
}

func (this *Solution) PickIndex() int {
	return sort.SearchInts(this.Prefix, rand.Int()%this.Sum)
}

func main() {

}
