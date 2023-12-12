package level2

import (
	. "AlgorithmCamp2023/utils"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	headA := BuildListNode([]interface{}{1, 2, 1, 4})
	headA = removeElements(headA, 1)
	PrintListNode(headA)
}
