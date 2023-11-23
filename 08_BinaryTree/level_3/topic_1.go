package level3

import . "AlgorithmCamp2023/utils"

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// 二叉树的最近公共祖先

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// 如果当前节点为目标节点，则直接返回，不用判断底下节点，就算底下节点为目标节点，当前节点也是目标节点的公共祖先
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 当前节点为公共祖先
	if left != nil && right != nil {
		return root
	}

	// 把公共祖先传到最外层
	if left != nil {
		return left
	}
	return right
}
