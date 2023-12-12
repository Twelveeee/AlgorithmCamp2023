package level3

import (
	"container/heap"
)

// 数据流的中位数
// https://leetcode.cn/problems/find-median-from-data-stream/description/
type MedianFinder struct {
	maxHeap *MaxHeap
	minHeap *MinHeap
}

func Constructor() MedianFinder {
	maxHeap, minHeap := &MaxHeap{}, &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)

	return MedianFinder{maxHeap, minHeap}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.minHeap, num)
	heap.Push(this.maxHeap, heap.Pop(this.minHeap))

	for this.minHeap.Len() < this.maxHeap.Len() {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() > this.maxHeap.Len() {
		return float64(this.minHeap.Top())
	}
	return float64(this.minHeap.Top()+this.maxHeap.Top()) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type MinHeap []int

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Push(val any)       { (*h) = append((*h), val.(int)) }
func (h *MinHeap) Pop() any {
	val := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return val
}
func (h *MinHeap) Top() int { return (*h)[0] }

type MaxHeap []int

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Push(val any)       { (*h) = append((*h), val.(int)) }
func (h *MaxHeap) Pop() any           { val := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return val }
func (h *MaxHeap) Top() int           { return (*h)[0] }
