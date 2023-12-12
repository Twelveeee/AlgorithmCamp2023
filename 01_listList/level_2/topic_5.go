package level2

import (
	. "AlgorithmCamp2023/utils"
)

// 删除特定节点
// https://leetcode.cn/problems/remove-linked-list-elements/description/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	for head != nil && head.Val == val {
		head = head.Next
	}

	tail := head
	pre := &ListNode{}
	pre.Next = tail
	for tail != nil {
		if tail.Val == val {
			pre.Next = tail.Next
			tail = pre
		}

		pre = tail
		tail = tail.Next
	}

	return head
}
