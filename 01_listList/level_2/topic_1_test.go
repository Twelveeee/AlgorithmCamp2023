package level2

import (
	. "AlgorithmCamp2023/utils"
	"fmt"
	"testing"
)

func TestFindFirstCommonNode(t *testing.T) {
	headA := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}

	node3.Next = node4
	node2.Next = node3
	headA.Next = node2

	PrintListNode(headA)
	fmt.Println()

	headB := &ListNode{0, nil}
	node5 := &ListNode{5, nil}

	node5.Next = node3
	headB.Next = node5
	PrintListNode(headB)
	fmt.Println()

	firstCommonNode := findFirstCommonNode(headA, headB)
	fmt.Println(firstCommonNode, node3)
}
