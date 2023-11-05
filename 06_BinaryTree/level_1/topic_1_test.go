package level1

import (
	"testing"
)

func makeTreeNode() *TreeNode {
	node1 := &TreeNode{1, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node4 := &TreeNode{4, nil, nil}
	node5 := &TreeNode{5, nil, nil}
	node6 := &TreeNode{6, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node8 := &TreeNode{8, nil, nil}

	node4.Left = node1
	node4.Right = node2
	node6.Left = node7
	node6.Right = node8

	node5.Left = node4
	node5.Right = node6
	return node5
}

func TestPreOrderTraversal(t *testing.T) {
	root := makeTreeNode()
	root.PreOrderTraversal()
	// 5412678
}

func TestInOrderTraversal(t *testing.T) {
	root := makeTreeNode()
	root.InOrderTraversal()
	// 1425768
}

func TestPostOrderTraversal(t *testing.T) {
	root := makeTreeNode()
	root.PostOrderTraversal()
	// 1247865
}
