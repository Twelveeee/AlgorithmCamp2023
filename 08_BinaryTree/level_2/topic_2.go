package level2

import (
	. "AlgorithmCamp2023/utils"
)

// 高度平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHight := height(node.Left)
	rightHight := height(node.Right)

	if leftHight == -1 || rightHight == -1 || abs(leftHight-rightHight) > 1 {
		return -1
	}

	return maxV2(leftHight, rightHight) + 1
}

func maxV2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
