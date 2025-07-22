package heap_pq

import "fmt"

type HeapScratch struct {
	arr []int
}

// getParent gets the parent index of current index
func (h *HeapScratch) getParent(idx int) int {
	return (idx - 1) / 2
}

func (h *HeapScratch) getLeft(idx int) int {
	return (idx * 2) + 1
}

func (h *HeapScratch) getRight(idx int) int {
	return (idx * 2) + 2
}

func (h *HeapScratch) Len() int {
	return len(h.arr)
}

func (h *HeapScratch) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

// Push pushes the element in heap time complexity O(logN) where n is the current size of heap
func (h *HeapScratch) Push(x int) {
	h.arr = append(h.arr, x)
	idx := len(h.arr) - 1
	for idx > 0 && h.arr[h.getParent(idx)] < h.arr[idx] {
		h.swap(h.getParent(idx), idx)
		idx = h.getParent(idx)
	}
}

func (h *HeapScratch) Heapify(idx int) {
	largest := idx

	if h.getLeft(idx) < h.Len() && h.arr[largest] < h.arr[h.getLeft(idx)] {
		largest = h.getLeft(idx)
	}
	if h.getRight(idx) < h.Len() && h.arr[largest] < h.arr[h.getRight(idx)] {
		largest = h.getRight(idx)
	}

	if largest != idx {
		h.swap(idx, largest)
		h.Heapify(idx)
	}

}

func (h *HeapScratch) Pop() int {
	n := len(h.arr)
	h.swap(0, n-1)
	x := h.arr[n-1]
	h.arr = h.arr[:n-1]
	h.Heapify(0)
	return x
}

func (h *HeapScratch) IsEmpty() bool {
	return h.Len() == 0
}

func (h *HeapScratch) Peek() int {
	if h.IsEmpty() {
		return -1
	}
	return h.arr[0]
}

func NewHeapScratch() *HeapScratch {
	return &HeapScratch{}
}

func TestHeapScratch() {
	h := NewHeapScratch()
	h.Push(8)
	h.Push(9)
	h.Push(13)
	h.Push(14)
	h.Push(15)

	for !h.IsEmpty() {
		fmt.Println(h.Pop())
	}

}
