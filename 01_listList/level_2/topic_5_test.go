package level2

import (
	"testing"
)

func TestRemoveElements(t *testing.T) {
	headA := InitLinkNodeByList([]int{1, 2, 1, 4})
	headA = removeElements(headA, 1)
	headA.PrintList()
}
