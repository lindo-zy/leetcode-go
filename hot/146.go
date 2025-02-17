package main

func main() {

}

type Node struct {
	Key, Val   int
	Prev, Next *Node
}
type LRUCache struct {
	Capacity int
	Hash     map[int]*Node
	Head     *Node
	Tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Capacity: capacity,
		Hash:     make(map[int]*Node, capacity),
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	return 0
}

func (this *LRUCache) Put(key int, value int) {

}
