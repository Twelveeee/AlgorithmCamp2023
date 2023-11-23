package level2

type Node struct {
	Val      int
	Children []*Node
}

// N 叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/description/

func maxDepthV3(root *Node) int {
	if root == nil {
		return 0
	}

	maxDep := 0

	for _, node := range root.Children {
		maxDep = maxV3(maxDep, maxDepthV3(node))
	}

	return maxDep + 1
}

func maxV3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
