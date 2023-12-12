package utils

import (
	"testing"
)

func TestBuilListNode(t *testing.T) {
	node1 := BuildListNode([]interface{}{1, 2, 3, 4})
	PrintListNode(node1)

	node2 := BuildListNode([]interface{}{1, 2, 2, 3, 3})
	PrintListNode(node2)
}
