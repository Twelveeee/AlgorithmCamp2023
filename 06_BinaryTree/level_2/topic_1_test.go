package level2

import (
	"fmt"
	"testing"

	. "AlgorithmCamp2023/utils"
)

func makeTreeNodeV1() *TreeNode {
	node3 := &TreeNode{3, nil, nil}
	node9 := &TreeNode{9, nil, nil}
	node20 := &TreeNode{20, nil, nil}
	node15 := &TreeNode{15, nil, nil}
	node7 := &TreeNode{7, nil, nil}

	node20.Left = node15
	node20.Right = node7

	node3.Left = node9
	node3.Right = node20
	return node3
}

func TestLevelOrder(t *testing.T) {
	root := makeTreeNodeV1()
	ans := levelOrder(root)
	for _, list := range ans {
		for _, v := range list {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
