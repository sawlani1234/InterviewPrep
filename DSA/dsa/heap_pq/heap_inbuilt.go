package heap_pq

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func TestHeap() {
	h := &IntHeap{}

	heap.Init(h)

	heap.Push(h, 9)
	heap.Push(h, 8)
	heap.Push(h, 10)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}

}
