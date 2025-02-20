package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

// 定义一个最小堆
type IntHeap [][2]int // 堆中的元素为 [频率, 数字]

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] } // 最小堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	res := make([]int, k)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	h := &IntHeap{}
	heap.Init(h)

	for num, fq := range freq {
		heap.Push(h, [2]int{fq, num})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[1]
	}

	return res
}
