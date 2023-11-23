package level3

import (
	. "AlgorithmCamp2023/utils"
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := BuildTree([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
	node5 := &TreeNode{5, nil, nil}
	node1 := &TreeNode{1, nil, nil}
	node4 := &TreeNode{4, nil, nil}
	node6 := &TreeNode{6, nil, nil}

	fmt.Println(lowestCommonAncestor(root, node5, node1))
	fmt.Println(lowestCommonAncestor(root, node5, node4))
	fmt.Println(lowestCommonAncestor(root, node4, node6))

}
