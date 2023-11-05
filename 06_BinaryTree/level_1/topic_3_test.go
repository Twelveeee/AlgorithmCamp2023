package level1

import (
	"testing"
)

func TestIteratePreOrderTraversal(t *testing.T) {
	root := makeTreeNode()
	root.IteratePreOrderTraversal()
	// 5412678
}

func TestIteratePreOrderTraversalV2(t *testing.T) {
	root := makeTreeNode()
	root.IteratePreOrderTraversalV2()
	// 5412678
}

func TestIterateInOrderTraversal(t *testing.T) {
	root := makeTreeNode()
	root.IterateInOrderTraversal()
	// 1425768
}

func TestIteratePostOrderTraversal(t *testing.T) {
	root := makeTreeNode()
	root.IteratePostOrderTraversal()
	// 8762145
}
