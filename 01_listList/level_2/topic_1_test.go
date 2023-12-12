package level2

import (
	. "AlgorithmCamp2023/utils"
	"fmt"
	"testing"
)

func TestFindFirstCommonNode(t *testing.T) {
	headA := InitLinkNode(1)
	node2 := InitLinkNode(2)
	node3 := InitLinkNode(3)
	node4 := InitLinkNode(4)

	node3.Next = node4
	node2.Next = node3
	headA.Next = node2

	PrintListNode(headA)
	fmt.Println()

	headB := InitLinkNode(0)
	node5 := InitLinkNode(5)

	node5.Next = node3
	headB.Next = node5
	PrintListNode(headB)
	fmt.Println()

	firstCommonNode := findFirstCommonNode(headA, headB)
	fmt.Println(firstCommonNode, node3)
}
