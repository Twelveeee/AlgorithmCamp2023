package level2

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	headA := InitLinkNodeByList([]int{2, 4, 6})
	headB := InitLinkNodeByList([]int{1, 3, 4})

	headC := mergeTwoLists(headA, headB)
	headC.PrintList()
}
