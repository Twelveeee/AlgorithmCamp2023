package level1

import (
	. "AlgorithmCamp2023/utils"
)

// 二叉树的路径总和
// https://leetcode.cn/problems/path-sum/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val-targetSum == 0
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
