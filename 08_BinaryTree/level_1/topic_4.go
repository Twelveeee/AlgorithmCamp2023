package level1

import (
	. "AlgorithmCamp2023/utils"
	"strconv"
)

// 二叉树的所有路径
// https://leetcode.cn/problems/binary-tree-paths/description/

var pathList []string

func binaryTreePaths(root *TreeNode) []string {
	pathList = make([]string, 0)
	dfs(root, "")
	return pathList
}

func dfs(node *TreeNode, path string) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		pathList = append(pathList, path+strconv.Itoa(node.Val))
		return
	}

	dfs(node.Left, path+strconv.Itoa(node.Val)+"->")
	dfs(node.Right, path+strconv.Itoa(node.Val)+"->")
}
