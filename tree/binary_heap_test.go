package tree

import (
	"log"
	"testing"
)

func TestMinBinaryHeap(t *testing.T) {

	heap := NeCBTree(10)
	heap.Push(4)
	heap.Push(3)
	heap.Push(1)
	heap.Push(7)
	heap.Push(5)
	heap.Push(8)
	log.Println(heap)
	log.Println(heap.Peek())

	log.Println(heap.Pop())
	log.Println(heap)

}
