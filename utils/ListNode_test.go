package utils

import (
	"fmt"
	"testing"
)

func printListNode(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	printListNode(head.Next)
}

func TestBuilListNode(t *testing.T) {
	node1 := BuildListNode([]interface{}{1, 2, 3, 4})
	printListNode(node1)

	node2 := BuildListNode([]interface{}{1, 2, 2, 3, 3})
	printListNode(node2)
}
