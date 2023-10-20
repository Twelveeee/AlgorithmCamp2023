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
