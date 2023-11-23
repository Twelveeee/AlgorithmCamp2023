package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(list []interface{}) *TreeNode {
	if len(list) == 0 {
		return nil
	}

	queue := make([]*TreeNode, 0)
	index := 0
	head := &TreeNode{list[index].(int), nil, nil}

	queue = append(queue, head)
	for index < len(list)-1 {
		index++
		parent := queue[0]
		queue = queue[1:]
		if index < len(list) && list[index] != nil {
			node := &TreeNode{list[index].(int), nil, nil}
			parent.Left = node
			queue = append(queue, node)
		}

		index++
		if index < len(list) && list[index] != nil {
			node := &TreeNode{list[index].(int), nil, nil}
			parent.Right = node
			queue = append(queue, node)
		}
	}

	return head
}
