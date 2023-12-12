package level1

import (
	"fmt"
	"testing"
)

func TestReverseListV1(t *testing.T) {
	node := InitLinkNodeByList([]int{1, 2, 3, 4})
	node.PrintList()

	node = reverseListV1(node)
	fmt.Println("**after**")

	node.PrintList()
}

func TestReverseListV2(t *testing.T) {
	node := InitLinkNodeByList([]int{1, 2, 3, 4})
	node.PrintList()

	node = reverseListV2(node)
	fmt.Println("**after**")

	node.PrintList()
}

func TestReverseListV3(t *testing.T) {
	node := InitLinkNodeByList([]int{1, 2, 3, 4})
	node.PrintList()

	node = reverseListV3(node)
	fmt.Println("**after**")

	node.PrintList()
}

func TestReverseListV4(t *testing.T) {
	node := InitLinkNodeByList([]int{1, 2, 3, 4})
	node.PrintList()

	node = reverseListV4(node)
	fmt.Println("**after**")

	node.PrintList()
}

// 虚拟头节点
func reverseListV4(node *LinkNode) *LinkNode {
	dummyHeadNode := &LinkNode{}
	dummyHeadNode.Next = node

	pre, curr, temp := dummyHeadNode, dummyHeadNode.Next, &LinkNode{}
	for curr.Next != nil {
		temp = curr.Next
		curr.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}

	return dummyHeadNode.Next
}

func TestReverseListV5(t *testing.T) {
	node := InitLinkNodeByList([]int{1, 2, 3, 4})
	node.PrintList()

	node = reverseListV5(node)
	fmt.Println("**after**")

	node.PrintList()
}

// 递归
func reverseListV5(node *LinkNode) *LinkNode {
	if node.Next == nil || node == nil {
		return node
	}

	tail := reverseListV5(node.Next)

	node.Next.Next = node
	node.Next = nil

	return tail
}
