package level3

import (
	"fmt"
	"testing"

	. "AlgorithmCamp2023/utils"
)

func TestReverseKGroup(t *testing.T) {
	node := BuildListNode([]interface{}{1, 2, 3, 4, 5})
	// node := BuildListNode([]interface{}{1, 2})
	PrintListNode(node)

	node = reverseKGroup(node, 2)
	fmt.Println("**after**")

	PrintListNode(node)
}
