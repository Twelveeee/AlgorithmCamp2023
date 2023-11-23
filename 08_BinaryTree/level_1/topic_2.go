package level1

import . "AlgorithmCamp2023/utils"

// https://leetcode.cn/problems/symmetric-tree/description/
// 对称二叉树

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return subIsSymmetric(root.Left, root.Right)
}

func subIsSymmetric(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	return subIsSymmetric(node1.Left, node2.Right) && subIsSymmetric(node1.Right, node2.Left)
}
