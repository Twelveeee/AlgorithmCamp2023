package level2

import (
	. "AlgorithmCamp2023/utils"
	"container/heap"
)

// 合并 K 个升序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListNodeHeap{}
	heap.Init(h)

	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	dummyHead := &ListNode{}
	curr := dummyHead

	for len(*h) > 0 {
		node := heap.Pop(h).(*ListNode)
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		curr.Next = node
		curr = curr.Next
	}

	return dummyHead.Next
}

type ListNodeHeap []*ListNode

func (h *ListNodeHeap) Len() int           { return len(*h) }
func (h *ListNodeHeap) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val }
func (h *ListNodeHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *ListNodeHeap) Push(val any)       { *h = append(*h, val.(*ListNode)) }
func (h *ListNodeHeap) Pop() any           { x := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return x }
