package learn_golang

import (
	"container/heap"
	"fmt"
	"testing"
)

type intHeap []int

// Len is the number of elements in the collection.
func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements with indexes i and j.
func (h intHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	old := *h
	length := len(old)

	e := old[length-1]

	*h = old[0 : length-1]

	return e
}

func TestHeap(t *testing.T) {

	s := intHeap{5, 2, 54, 76, 85, 2, 4}
	heap.Init(&s)

	for s.Len() > 0 {
		fmt.Println(heap.Pop(&s))
	}

}
