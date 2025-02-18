package main

import "fmt"
import (
	"container/heap"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

// 定义最小堆结构
type MinHeap1 []int

func (h MinHeap1) Len() int           { return len(h) }
func (h MinHeap1) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap1) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap1) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap1{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}
