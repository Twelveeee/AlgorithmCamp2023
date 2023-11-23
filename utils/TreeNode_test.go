package utils

import (
	"fmt"
	"testing"
)

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func TestBuildTree(t *testing.T) {
	root1 := BuildTree([]interface{}{1, 2, 3, 4})
	printTree(root1)

	root2 := BuildTree([]interface{}{1, 2, 2, nil, 3, nil, 3})
	printTree(root2)
}
