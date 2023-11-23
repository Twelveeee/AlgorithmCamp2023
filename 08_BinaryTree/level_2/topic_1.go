package level2

import . "AlgorithmCamp2023/utils"

// 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/

func maxDepth(root *TreeNode) int {
	return subMaxDepth(root, 1)
}

func subMaxDepth(root *TreeNode, nowDepth int) int {
	if root == nil {
		return nowDepth - 1
	}

	if root.Left == nil && root.Right == nil {
		return nowDepth
	}

	return max(subMaxDepth(root.Right, nowDepth+1), subMaxDepth(root.Left, nowDepth+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
