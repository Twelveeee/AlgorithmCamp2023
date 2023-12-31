package level2

import (
	. "AlgorithmCamp2023/utils"
)

// 判断链表是否为回文序列
// https://leetcode.cn/problems/palindrome-linked-list/description/
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	half := slow.Next
	var reverseNode, curr, next *ListNode
	curr = half
	for curr != nil {
		next = curr.Next
		curr.Next = reverseNode
		reverseNode = curr
		curr = next
	}

	for reverseNode != nil {
		if reverseNode.Val != head.Val {
			return false
		}
		reverseNode = reverseNode.Next
		head = head.Next
	}

	return true
}
