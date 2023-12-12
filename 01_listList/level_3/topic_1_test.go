package level3

import (
	. "AlgorithmCamp2023/utils"
	"fmt"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	headA := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}
	node5 := &ListNode{5, nil}
	node6 := &ListNode{6, nil}
	node7 := &ListNode{7, nil}

	headA.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node3

	// headA.PrintList()
	fmt.Println()

	firstCommonNode := detectCycle(headA)
	fmt.Println(firstCommonNode.Val)
}
