package level3

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	node := InitListNodeByList([]int{1, 2, 3, 4, 5})
	// node := InitListNodeByList([]int{1, 2})
	node.PrintList()

	node = reverseKGroup(node, 2)
	fmt.Println("**after**")

	node.PrintList()
}
