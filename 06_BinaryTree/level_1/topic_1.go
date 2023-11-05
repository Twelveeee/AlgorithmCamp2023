package level1

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func (root *TreeNode) PreOrderTraversal() {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	root.Left.PreOrderTraversal()
	root.Right.PreOrderTraversal()
}

// 中序遍历
func (root *TreeNode) InOrderTraversal() {
	if root == nil {
		return
	}
	root.Left.InOrderTraversal()
	fmt.Println(root.Val)
	root.Right.InOrderTraversal()
}

// 后序遍历
func (root *TreeNode) PostOrderTraversal() {
	if root == nil {
		return
	}
	root.Right.PostOrderTraversal()
	root.Left.PostOrderTraversal()
	fmt.Println(root.Val)
}
