package level3

import (
	. "AlgorithmCamp2023/utils"
)

// K 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummpHead, tempNext := &ListNode{}, &ListNode{}
	dummpHead.Next = head

	pre := dummpHead
	curr := pre.Next
	for curr != nil {
		// check
		tail := curr
		for i := 0; i < k; i++ {
			if tail == nil {
				return dummpHead.Next
			}
			tail = tail.Next
		}

		// do reverse
		for i := 0; i < k-1; i++ {
			tempNext = curr.Next
			curr.Next = tempNext.Next
			tempNext.Next = pre.Next
			pre.Next = tempNext
		}

		pre = curr
		curr = curr.Next

	}

	return dummpHead.Next
}
