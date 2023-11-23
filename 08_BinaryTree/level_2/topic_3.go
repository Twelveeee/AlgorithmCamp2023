package level2

import (
	. "AlgorithmCamp2023/utils"
)

// 二叉树的最小深度
// https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	minDep := 10000

	if root.Left != nil {
		minDep = min(minDep, minDepth(root.Left))
	}
	if root.Right != nil {
		minDep = min(minDep, minDepth(root.Right))
	}

	return minDep + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
