package level2

import (
	. "AlgorithmCamp2023/utils"
)

// https://leetcode.cn/problems/add-two-numbers-ii/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := make([]int, 0, 100)
	list2 := make([]int, 0, 100)

	for l1 != nil {
		list1 = append(list1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		list2 = append(list2, l2.Val)
		l2 = l2.Next
	}

	ans := &ListNode{}

	index1, index2 := len(list1)-1, len(list2)-1
	carry := 0
	for index1 >= 0 || index2 >= 0 || carry != 0 {
		temp := 0
		if carry != 0 {
			temp += carry
			carry = 0
		}
		if index1 >= 0 {
			temp += list1[index1]
			index1--
		}
		if index2 >= 0 {
			temp += list2[index2]
			index2--
		}

		if temp >= 10 {
			carry, temp = temp/10, temp%10
		}

		ans.Next = &ListNode{temp, ans.Next}

	}

	return ans.Next
}
