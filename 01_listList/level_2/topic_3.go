package level2

import (
	. "AlgorithmCamp2023/utils"
)

// 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/description/
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	tail := &ListNode{0, nil}
	head := tail

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			tail = tail.Next
			list1 = list1.Next
			continue
		}
		tail.Next = list2
		tail = tail.Next
		list2 = list2.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return head.Next
}
