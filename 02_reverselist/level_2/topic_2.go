package level2

import (
	. "AlgorithmCamp2023/utils"
)

// https://leetcode.cn/problems/swap-nodes-in-pairs/description/
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	pre, curr, next := dummyHead, dummyHead.Next, &ListNode{}

	for curr != nil && curr.Next != nil {
		next = curr.Next
		curr.Next = curr.Next.Next
		pre.Next = next
		next.Next = curr

		pre = pre.Next.Next
		curr = pre.Next
	}

	return dummyHead.Next
}
