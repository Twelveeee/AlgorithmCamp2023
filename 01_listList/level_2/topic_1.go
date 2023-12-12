package level2

import (
	. "AlgorithmCamp2023/utils"
)

// 两个链表第一个公共子节点 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/
func findFirstCommonNode(headA, headB *ListNode) *ListNode {

	lenA, lenB := 0, 0

	nodeA, nodeB := headA, headB
	for nodeA != nil {
		nodeA = nodeA.Next
		lenA++
	}

	for nodeB != nil {
		nodeB = nodeB.Next
		lenB++
	}

	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}

	} else {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
