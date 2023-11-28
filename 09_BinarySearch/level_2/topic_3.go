package level2

import (
	. "AlgorithmCamp2023/utils"
	"math"
)

// 验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/description/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	low, high := math.MinInt, math.MaxInt
	return subValidBST(root.Left, low, root.Val) && subValidBST(root.Right, root.Val, high)
}

func subValidBST(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}

	if root.Val >= high || root.Val <= low {
		return false
	}

	return subValidBST(root.Left, low, root.Val) && subValidBST(root.Right, root.Val, high)
}
