package level1

import (
	. "AlgorithmCamp2023/utils"
	"fmt"
	"testing"
)

func TestReverseListV1(t *testing.T) {
	node := BuildListNode([]interface{}{1, 2, 3, 4})
	PrintListNode(node)

	node = reverseListV1(node)
	fmt.Println("**after**")

	PrintListNode(node)
}

func TestReverseListV2(t *testing.T) {
	node := BuildListNode([]interface{}{1, 2, 3, 4})
	PrintListNode(node)

	node = reverseListV2(node)
	fmt.Println("**after**")

	PrintListNode(node)
}

func TestReverseListV3(t *testing.T) {
	node := BuildListNode([]interface{}{1, 2, 3, 4})
	PrintListNode(node)

	node = reverseListV3(node)
	fmt.Println("**after**")

	PrintListNode(node)
}

func TestReverseListV4(t *testing.T) {
	node := BuildListNode([]interface{}{1, 2, 3, 4})
	PrintListNode(node)

	node = reverseListV4(node)
	fmt.Println("**after**")

	PrintListNode(node)
}

// 虚拟头节点
func reverseListV4(node *ListNode) *ListNode {
	dummyHeadNode := &ListNode{}
	dummyHeadNode.Next = node

	pre, curr, temp := dummyHeadNode, dummyHeadNode.Next, &ListNode{}
	for curr.Next != nil {
		temp = curr.Next
		curr.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}

	return dummyHeadNode.Next
}

func TestReverseListV5(t *testing.T) {
	node := BuildListNode([]interface{}{1, 2, 3, 4})
	PrintListNode(node)

	node = reverseListV5(node)
	fmt.Println("**after**")

	PrintListNode(node)
}

// 递归
func reverseListV5(node *ListNode) *ListNode {
	if node.Next == nil || node == nil {
		return node
	}

	tail := reverseListV5(node.Next)

	node.Next.Next = node
	node.Next = nil

	return tail
}
