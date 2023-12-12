package level2

import (
	"math"

	. "AlgorithmCamp2023/utils"
)

// 在每个树行中找最大值
// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/description/

func largestValues(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	list := []*TreeNode{root}
	for len(list) > 0 {
		maxVal := math.MinInt
		tempList := []*TreeNode{}
		for _, v := range list {
			if maxVal < v.Val {
				maxVal = v.Val
			}

			if v.Left != nil {
				tempList = append(tempList, v.Left)
			}
			if v.Right != nil {
				tempList = append(tempList, v.Right)
			}
		}
		ans = append(ans, maxVal)
		list = tempList
	}
	return ans
}
