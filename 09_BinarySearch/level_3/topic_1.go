package level3

import . "AlgorithmCamp2023/utils"

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/
// 将有序数组转换为二叉搜索树

func sortedArrayToBST(nums []int) *TreeNode {
	return subSortedArrayToBST(nums, 0, len(nums)-1)
}

func subSortedArrayToBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = subSortedArrayToBST(nums, left, mid-1)
	root.Right = subSortedArrayToBST(nums, mid+1, right)
	return root
}

func sortedArrayToBSTV2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = sortedArrayToBSTV2(nums[:mid])
	root.Right = sortedArrayToBSTV2(nums[mid+1:])

	return root
}
