package level2

import (
	"testing"

	. "AlgorithmCamp2023/utils"
)

func TestMergeTwoLists(t *testing.T) {
	headA := BuildListNode([]interface{}{2, 4, 6})
	headB := BuildListNode([]interface{}{1, 3, 4})

	headC := mergeTwoLists(headA, headB)
	PrintListNode(headC)
}
