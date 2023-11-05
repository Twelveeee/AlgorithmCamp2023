package level2

// 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/description/

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	level := 0
	list := []*TreeNode{root}
	for len(list) > 0 {
		tempList := []*TreeNode{}
		ans = append(ans, []int{})
		for _, v := range list {
			ans[level] = append(ans[level], v.Val)
			if v.Left != nil {
				tempList = append(tempList, v.Left)
			}
			if v.Right != nil {
				tempList = append(tempList, v.Right)
			}
		}
		level++
		list = tempList
	}

	return ans
}
