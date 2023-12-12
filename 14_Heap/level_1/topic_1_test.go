package level1

import (
	"container/heap"
	"testing"
)

func TestGoHeap(t *testing.T) {
	h := newIntHeap()

	heap.Push(h, 5)
	heap.Push(h, 4)
	heap.Push(h, 1)
	heap.Push(h, 3)
	heap.Push(h, 2)

	if h.Len() != 5 {
		t.Error("Heap.Len() != 5")
	}

	for h.Len() > 0 {
		t.Log(heap.Pop(h))
	}

}
