package level1

import (
	"fmt"
)

// 迭代法实现二叉树遍历
// 前序遍历
func (root *TreeNode) IteratePreOrderTraversal() {
	if root == nil {
		return
	}

	ans := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 前序
		ans = append(ans, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	fmt.Println(ans)
}

// 前序遍历
func (root *TreeNode) IteratePreOrderTraversalV2() {
	if root == nil {
		return
	}

	ans := []int{}
	stack := []*TreeNode{}
	curr := root
	for len(stack) > 0 || curr != nil {
		for curr != nil {
			ans = append(ans, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr = curr.Right
	}

	fmt.Println(ans)
}

// 中序遍历
func (root *TreeNode) IterateInOrderTraversal() {
	if root == nil {
		return
	}
	ans := []int{}
	stack := []*TreeNode{}
	curr := root
	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, curr.Val)
		curr = curr.Right
	}

	fmt.Println(ans)
}

// 后序遍历
func (root *TreeNode) IteratePostOrderTraversal() {
	if root == nil {
		return
	}

	ans := []int{}
	stack := []*TreeNode{}
	curr := root
	for len(stack) > 0 || curr != nil {
		for curr != nil {
			ans = append(ans, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr = curr.Right
	}

	// 前序遍历之后翻转一下就是后序
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	fmt.Println(ans)
}
