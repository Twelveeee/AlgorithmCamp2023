package level3

import (
	. "AlgorithmCamp2023/utils"
	"fmt"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	headA := InitLinkNode(1)
	node2 := InitLinkNode(2)
	node3 := InitLinkNode(3)
	node4 := InitLinkNode(4)
	node5 := InitLinkNode(5)
	node6 := InitLinkNode(6)
	node7 := InitLinkNode(7)

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
