package level1

import "container/heap"

// 基于内部包实现栈
type StackV3 []int

func InitStackV3() *StackV3 {
	s := new(StackV3)
	heap.Init(s)
	return s
}

func (h *StackV3) Len() int {
	return len(*h)
}

func (h *StackV3) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *StackV3) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *StackV3) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *StackV3) Pop() interface{} {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}
