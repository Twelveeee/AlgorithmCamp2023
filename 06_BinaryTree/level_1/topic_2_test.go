package level1

import (
	"testing"
)

func TestBuildTree(t *testing.T) {

	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)

	// fmt.Println(root)
	root.PreOrderTraversal()
}
