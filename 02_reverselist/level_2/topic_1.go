package level2

import (
	. "AlgorithmCamp2023/utils"
)

// 指定区间翻转
// https://leetcode.cn/problems/reverse-linked-list-ii/description/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	left = left - 1
	right = right - 1

	dummyHead := &ListNode{}
	dummyHead.Next = head

	pre := dummyHead
	for i := 0; i < left; i++ {
		pre = pre.Next
	}

	curr := pre.Next
	temp := &ListNode{}
	for i := 0; i < right-left; i++ {
		temp = curr.Next
		curr.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}

	return dummyHead.Next
}
